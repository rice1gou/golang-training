package user

import (
	"fmt"
	"regexp"
)

type User struct {
	ID    string
	Age   int
	Name  string
	Email string
}

func NewUser(id string, age int, name, email string) *User {
	return &User{ID: id, Age: age, Name: name, Email: email}
}

func InitUser(re *regexp.Regexp, u *User) error {
	if match := re.MatchString(u.ID); !match {
		return fmt.Errorf("invalid identity")
	}
	fmt.Println("New User Created!")
	fmt.Printf("id: %v, name: %v, age: %v, email: %v\n", u.ID, u.Name, u.Age, u.Email)
	return nil
}
