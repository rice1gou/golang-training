package user

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type User struct {
	UserOid  string
	UserId   string
	UserName string
	Password string
}

type userRepository struct {
	db *sql.DB
}

type UserRepository interface {
	CreateUserTable() error
	SaveUser(u *User) error
	FetchUsers() ([]*User, error)
	FetchUserDetails(useroid string) (*User, error)
	ModifyUser(uoid, ui, un string) error
	DeleteUser(useroid string) error
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func NewUser(userid, username, password string) *User {
	uuidWithHyphen, _ := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return &User{UserOid: uuid, UserId: userid, UserName: username, Password: password}
}

func (ur *userRepository) CreateUserTable() error {
	sqlStr := `CREATE TABLE IF NOT EXISTS m_users(
		useroid  VARCHAR(128),
		userid VARCHAR(64) NOT NULL,
		username VARCHAR(16) NOT NULL,
		password VARCHAR(16) NOT NULL
	);`
	_, err := ur.db.Exec(sqlStr)
	if err != nil {
		return err
	}
	fmt.Println("Success for Create User Table")
	return nil
}

func (ur *userRepository) SaveUser(u *User) error {
	sqlStr := "INSERT INTO m_users(useroid, userid, username, password) VALUES($1,$2,$3,$4);"
	_, err := ur.db.Exec(sqlStr, u.UserOid, u.UserId, u.UserName, u.Password)
	if err != nil {
		return fmt.Errorf("save user: %w", err)
	}
	return nil
}

func (ur *userRepository) FetchUsers() ([]*User, error) {
	sqlStr := "SELECT useroid, userid, username FROM m_users;"
	rows, err := ur.db.Query(sqlStr)
	if err != nil {
		return nil, fmt.Errorf("fetch users: %w", err)
	}
	defer rows.Close()
	var users []*User

	for rows.Next() {
		var u User
		err := rows.Scan(&u.UserOid, &u.UserId, &u.UserName)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		users = append(users, &u)
	}
	return users, nil
}

func (ur *userRepository) FetchUserDetails(useroid string) (*User, error) {
	sqlStr := "SELECT userid, username FROM m_users WHERE useroid=$1;"
	row := ur.db.QueryRow(sqlStr, useroid)
	var u User
	err := row.Scan(&u.UserId, &u.UserName)
	if err != nil {
		return &User{}, fmt.Errorf("scan: %w", err)
	}
	return &u, nil
}

func (ur *userRepository) ModifyUser(uoid, ui, un string) error {
	sqlStr := "UPDATE m_users SET userid=$2, username=3 WHERE useroid=$1;"
	_, err := ur.db.Exec(sqlStr, uoid, ui, un)
	if err != nil {
		return fmt.Errorf("update: %w", err)
	}
	return nil
}

func (ur *userRepository) DeleteUser(useroid string) error {
	sqlStr := "DELETE FROM m_users WHERE useroid=$1;"
	_, err := ur.db.Exec(sqlStr, useroid)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}
