package main

import (
	"errors"
	"flag"
	"fmt"
	"identity/pkg/user"
	"os"
	"regexp"
)

func main() {
	var (
		id    = flag.String("id", "0000", "register your identity")
		age   = flag.Int("age", 0, "register your identity")
		name  = flag.String("name", "", "register your identity")
		email = flag.String("email", "", "register your identity")
	)

	flag.Parse()

	re := regexp.MustCompile(`^\d{4}$`)
	u := user.NewUser(user.UserID(*id), *age, *name, *email)

	if err := initUser(re, u); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initUser(re *regexp.Regexp, u *user.User) error {
	if !u.ID.IsCorrectID(re) {
		return errors.New("invalid identity")
	}
	fmt.Println("New User Created!")
	fmt.Printf("id: %v, name: %v, age: %v, email: %v\n", u.ID, u.Name, u.Age, u.Email)
	return nil
}
