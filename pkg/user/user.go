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

func NewUser(userid, username, password string) *User {
	uuidWithHyphen, _ := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return &User{UserOid: uuid, UserId: userid, UserName: username, Password: password}
}

func CreateUserTable(db *sql.DB) error {
	sqlStr := `CREATE TABLE IF NOT EXISTS m_users(
		useroid  VARCHAR(128),
		userid VARCHAR(64) NOT NULL,
		username VARCHAR(16) NOT NULL,
		password VARCHAR(16) NOT NULL
	);`
	_, err := db.Exec(sqlStr)
	if err != nil {
		return err
	}
	fmt.Println("Success for Create User Table")
	return nil
}

func SaveUser(db *sql.DB, u User) error {
	sqlStr := "INSERT INTO m_users(useroid, userid, username, password) VALUES($1,$2,$3,$4);"
	_, err := db.Exec(sqlStr, u.UserOid, u.UserId, u.UserName, u.Password)
	if err != nil {
		return fmt.Errorf("save user: %w", err)
	}
	return nil
}

func FetchUsers(db *sql.DB) ([]*User, error) {
	sqlStr := "SELECT useroid, userid, username FROM m_users;"
	rows, err := db.Query(sqlStr)
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

func FetchUserDetails(db *sql.DB, useroid string) (*User, error) {
	sqlStr := "SELECT userid, username FROM m_users WHERE useroid=$1;"
	row := db.QueryRow(sqlStr, useroid)
	var u User
	err := row.Scan(&u.UserId, &u.UserName)
	if err != nil {
		return &User{}, fmt.Errorf("scan: %w", err)
	}
	return &u, nil
}

func ModifyUser(db *sql.DB, u *User) error {
	sqlStr := "UPDATE m_users SET userid=?, username=? WHERE useroid=$1;"
	_, err := db.Exec(sqlStr, u.UserId, u.UserName, u.UserOid)
	if err != nil {
		return fmt.Errorf("update: %w", err)
	}
	return nil
}

func DeleteUser(db *sql.DB, useroid string) error {
	sqlStr := "DELETE FROM m_users WHERE useroid=$1;"
	_, err := db.Exec(sqlStr, useroid)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}
