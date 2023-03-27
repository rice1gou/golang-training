package handler

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

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

var tur = newTestUserRepository(
	[]*user.User{
		user.NewUser("test1@email.com", "testuser1", "password1"),
		user.NewUser("test2@email.com", "testuser2", "password2"),
		user.NewUser("test3@email.com", "testuser3", "password3"),
	},
)

func TestFetchUsersHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	got := httptest.NewRecorder()

	type args struct {
		tur user.UserRepository
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		// TODO: Add test cases.
		{"test1", args{tur}, 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FetchUsersHandler(tt.args.tur)(got, req)
			fmt.Println(got.Body, got.Code)
		})
	}
}

func TestFetchUserDetailsHandler(t *testing.T) {
	db, _ := tur.FetchUsers()
	oid := db[0].UserOid

	req := httptest.NewRequest(http.MethodGet, "/user/"+oid+"/", nil)
	got := httptest.NewRecorder()
	ctx := context.WithValue(req.Context(), router.PathParamCtxKey{}, []string{oid})

	type args struct {
		tur user.UserRepository
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{tur}, 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FetchUserDetailsHandler(tt.args.tur)(got, req.WithContext(ctx))
			fmt.Println(got.Body)
			fmt.Println(got.Header())
			if !reflect.DeepEqual(got.Code, tt.want) {
				t.Errorf("FetchUserDetailsHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveUserHandler(t *testing.T) {
	formData := url.Values{}
	formData.Set("userid", "test4@email.com")
	formData.Set("username", "testuser4")
	formData.Set("password", "password4")
	reqBody := strings.NewReader(formData.Encode())
	req := httptest.NewRequest(http.MethodPost, "/user", reqBody)
	got := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	type args struct {
		tur user.UserRepository
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		// TODO: Add test cases.
		{"test1", args{tur}, 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveUserHandler(tt.args.tur)(got, req)
			fmt.Println(got.Body)
			FetchUsersHandler(tt.args.tur)(got, req)
			fmt.Println(got.Body)

		})
	}
}
