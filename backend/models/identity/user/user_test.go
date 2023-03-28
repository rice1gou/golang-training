package user

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestNewUser(t *testing.T) {
	type args struct {
		userid   string
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.userid, tt.args.username, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_SaveUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ur := NewUserRepository(db)
	u := NewUser("test1@test.com", "testuser1", "password1")
	mock.ExpectExec("INSERT INTO m_users").
		WithArgs(u.UserOid, u.UserId, u.UserName, u.Password).WillReturnResult(sqlmock.NewResult(0, 1))
	// WillReturnRows(sqlmock.NewRows([]string{"useroid", "userid", "username", "password"}).AddRow(id, name))

	if err = ur.SaveUser(u); err != nil {
		t.Errorf("was expecting an error, but there was none")
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
