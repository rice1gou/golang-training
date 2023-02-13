package user

import (
	"fmt"
	"regexp"
)

type User struct {
	ID    ID
	Age   int
	Name  string
	Email string
}

func NewUser(id ID, age int, name, email string) *User {
	return &User{
		ID:    id,
		Age:   age,
		Name:  name,
		Email: email,
	}
}

type UserID struct {
	id string
}

func (uid *UserID) IsCorrectId(re string) (bool, error) {
	match, err := regexp.MatchString(re, uid.id)
	if err != nil {
		return false, fmt.Errorf("match string err: %w", err)
	}
	return match, nil
}

type ID interface {
	IsCorrectID(re string) (bool, error)
}
