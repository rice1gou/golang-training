package main

import (
	"flag"
	"fmt"
	"regexp"

	"github.com/rice1gou/golang-training/user"
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
	u := user.NewUser(*id, *age, *name, *email)

	if err := initUser(re, u); err != nil {
		fmt.Println(err.Error())
	}
}

func initUser(re *regexp.Regexp, u *user.User) error {
	if match := re.MatchString(u.ID); !match {
		return fmt.Errorf("invalid identity")
	}
	fmt.Println("New User Created!")
	fmt.Printf("id: %v, name: %v, age: %v, email: %v\n", u.ID, u.Name, u.Age, u.Email)
	return nil
}
