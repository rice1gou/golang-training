package main

import (
	"database/sql"
	"log"

	"fmt"
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
	err := run()
	if err != nil {
		log.Fatal(err)
	}
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

	mux := router.NewRouter()
	mux.Add(http.MethodGet, "/", handler.IndexHandler(db))
	mux.Add(http.MethodPost, "/signin", handler.SigninHandler(db))
	mux.Add(http.MethodGet, "/signout", handler.SignoutHandler(db))
	mux.Add(http.MethodGet, "/signup", handler.SignupHandler(db))
	mux.Add(http.MethodGet, "/user", handler.FetchUsersHandler(db))
	mux.Add(http.MethodPost, "/user", handler.SaveUserHandler(db))
	mux.Add(http.MethodGet, "/user/([^/]+)", handler.FetchUserDetailsHandler(db))
	mux.Add(http.MethodPost, "/user/([^/]+)", handler.ModifyUserHandler(db))
	mux.Add(http.MethodDelete, "/user/([^/]+)", handler.DeleteUserHandler(db))

	err = http.ListenAndServe(":8080", mux);
	if err != nil {
		return err
	}
	return nil
}
