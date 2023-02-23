package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/rice1gou/golang-training/pkg/user"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintln(w, "Register New User")
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user/" {
		http.NotFound(w, r)
		return
	}

	id := r.FormValue("id")
	name := r.FormValue("name")
	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	email := r.FormValue("email")

	re := regexp.MustCompile(`^\d{4}$`)
	u := user.NewUser(id, age, name, email)

	fmt.Fprintln(w, "New User Created!")
	fmt.Fprintf(w, "id: %v, name: %v, age: %v, email: %v\n", u.ID, u.Name, u.Age, u.Email)

	user.InitUser(re, u)
}
