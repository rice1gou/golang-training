package router

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/rice1gou/golang-training/pkg/user"
)

type router struct {
	routes []route
}

type route struct {
	method    string
	path      *regexp.Regexp
	pathParam []string
	handler   http.HandlerFunc
}

func NewRouter() *router {
	return &router{}
}

func (r *router) Add(method string, pattern string, handler http.HandlerFunc) *router {
	newRoute := route{method, regexp.MustCompile("^" + pattern + "$"), nil, handler}
	r.routes = append(r.routes, newRoute)
	return r
}

type pathParamCtxKey struct{}

func isMatchPath(r *router, req *http.Request) []route {
	var matches []route
	for _, route := range r.routes {
		match := route.path.FindStringSubmatch(req.URL.Path)
		if 0 < len(match) {
			route.pathParam = match[:1]
			matches = append(matches, route)
		}
	}
	return matches
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var allow []string
	matches := isMatchPath(r, req)
	if len(matches) < 0 {
		http.NotFound(w, req)
		return
	}
	for _, route := range matches {
		if route.method == req.Method {
			ctx := context.WithValue(req.Context(), pathParamCtxKey{}, route.pathParam)
			route.handler(w, req.WithContext(ctx))
			return
		}
		allow = append(allow, route.method)
	}
	w.Header().Set("Allow", strings.Join(allow, ", "))
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

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

	if err := user.InitUser(re, u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
