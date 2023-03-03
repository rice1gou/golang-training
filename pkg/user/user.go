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

func createUserTable(db *sql.DB) error {
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
	return nil
}

func saveUser(db *sql.DB, u User) error {
	sqlStr := `INSERT INTO m_users(useroid, userid, username, password) VALUES(?,?,?,?);`
	_, err := db.Exec(sqlStr, u.UserOid, u.UserId, u.UserName, u.Password)
	if err != nil {
		return fmt.Errorf("save user: %w", err)
	}
	return nil
}

func fetchUsers(db *sql.DB) error {
	sqlStr := `SELECT * FROM m_users;`
	rows, err := db.
	return nil
}

func fetchUserDetails() error {
	return nil
}

func modifyUser() error {
	return nil
}

func deleteUser() error {
	return nil
}
