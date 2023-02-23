package main

import (
	"log"
	"net/http"

	"github.com/rice1gou/golang-training/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("/user/", handler.UserHandler)

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
