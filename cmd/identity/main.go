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
	mux.Add(http.MethodPost, "/login", handler.IndexHandler)
	mux.Add(http.MethodGet, "/logout", handler.IndexHandler)
	mux.Add(http.MethodGet, "/user", handler.IndexHandler)
	mux.Add(http.MethodPost, "/user", handler.IndexHandler)
	mux.Add(http.MethodGet, "/", handler.IndexHandler)

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
