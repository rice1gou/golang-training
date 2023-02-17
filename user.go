package main

type User struct {
	ID    string
	Age   int
	Name  string
	Email string
}

func NewUser(id string, age int, name, email string) *User {
	return &User{ID: id, Age: age, Name: name, Email: email}
}
