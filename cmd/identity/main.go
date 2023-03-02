package main

import (
	"log"
	"net/http"

	"github.com/rice1gou/golang-training/handler"
	"github.com/rice1gou/golang-training/pkg/router"
)

func main() {
	mux := router.NewRouter()
	mux.Add(http.MethodGet, "/", handler.IndexHandler)
	mux.Add(http.MethodPost, "/signin", handler.SigninHandler)
	mux.Add(http.MethodGet, "/signout", handler.SignoutHandler)
	mux.Add(http.MethodGet, "/signup", handler.SignupHandler)
	mux.Add(http.MethodGet, "/user", handler.FetchUsersHandler)
	mux.Add(http.MethodPost, "/user", handler.RegisterUserHandler)
	mux.Add(http.MethodGet, "/user/([^/]+)", handler.FetchUserHandler)
	mux.Add(http.MethodPost, "/user/([^/]+)", handler.ModifyUserHandler)
	mux.Add(http.MethodDelete, "/user/([^/]+)", handler.DeleteUserHandler)

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
