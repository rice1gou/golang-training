package main

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/rice1gou/golang-training/models/identity/user"
	"github.com/rice1gou/golang-training/pkg/router"
)

type testUserRepository struct {
	db []*user.User
}

func newTestUserRepository(db []*user.User) user.UserRepository {
	return &testUserRepository{db}
}

func (tur *testUserRepository) CreateUserTable() error {
	return nil
}

func (tur *testUserRepository) SaveUser(u *user.User) error {
	tur.db = append(tur.db, u)
	return nil
}

func (tur *testUserRepository) FetchUsers() ([]*user.User, error) {
	return tur.db, nil
}

func (tur *testUserRepository) FetchUserDetails(useroid string) (*user.User, error) {
	for _, user := range tur.db {
		if user.UserOid == useroid {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (tur *testUserRepository) ModifyUser(uoid, ui, un string) error {
	for _, user := range tur.db {
		if user.UserOid == uoid {
			user.UserId = ui
			user.UserName = un
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

func (tur *testUserRepository) DeleteUser(useroid string) error {
	for i, user := range tur.db {
		if user.UserOid == useroid {
			tur.db[i] = tur.db[len(tur.db)-1]
			tur.db = tur.db[:len(tur.db)-1]
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

var tur = testUserRepository{
	[]*user.User{
		user.NewUser("test1@email.com", "testuser1", "password1"),
		user.NewUser("test2@email.com", "testuser2", "password2"),
		user.NewUser("test3@email.com", "testuser3", "password3"),
	},
}

func Test_run(t *testing.T) {
	type args struct {
		ur  user.UserRepository
		mux *router.Router
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := run(tt.args.ur, tt.args.mux); (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
