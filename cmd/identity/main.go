package main

import (
	"database/sql"

	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/rice1gou/golang-training/handler"
	"github.com/rice1gou/golang-training/pkg/router"
)

const (
	HOST     = ""
	USER     = ""
	PASSWORD = ""
	DATABASE = ""
)

func main() {
	mux := router.NewRouter()
	mux.Add(http.MethodGet, "/", handler.IndexHandler)
	mux.Add(http.MethodPost, "/signin", handler.SigninHandler)
	mux.Add(http.MethodGet, "/signout", handler.SignoutHandler)
	mux.Add(http.MethodGet, "/signup", handler.SignupHandler)
	mux.Add(http.MethodGet, "/user", handler.FetchUsersHandler)
	mux.Add(http.MethodPost, "/user", handler.SaveUserHandler)
	mux.Add(http.MethodGet, "/user/([^/]+)", handler.FetchUserHandler)
	mux.Add(http.MethodPost, "/user/([^/]+)", handler.ModifyUserHandler)
	mux.Add(http.MethodDelete, "/user/([^/]+)", handler.DeleteUserHandler)

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func run() error {
	driverName := "postgres"
	connectStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", HOST, USER, PASSWORD, DATABASE)
	db, err := sql.Open(driverName, connectStr)
	if err != nil {
		return fmt.Errorf("DB Open: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func createUserTable(db *sql.DB) error {
	queryStr := ``
	_, err := db.Exec(queryStr)
	if err != nil {
		return fmt.Errorf("Create Table: %w", err)
	}
	return nil
}
