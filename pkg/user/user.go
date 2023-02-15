package user

import (
	"regexp"
)

type User struct {
	ID    UserID
	Age   int
	Name  string
	Email string
}

func NewUser(id UserID, age int, name, email string) *User {
	return &User{ID: id, Age: age, Name: name, Email: email}
}

type UserID string

func (uid *UserID) IsCorrectID(re *regexp.Regexp) bool {
	match := re.MatchString(string(*uid))
	return match
}

type ID interface {
	IsCorrectID(re string) (bool, error)
}
