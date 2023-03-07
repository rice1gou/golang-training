package main

import (
	"database/sql"
	"log"
	"os"

	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/rice1gou/golang-training/handler"
	"github.com/rice1gou/golang-training/pkg/router"
	"github.com/rice1gou/golang-training/pkg/user"
)

var (
	HOST     = os.Getenv("DB_HOST_NAME")
	USER     = os.Getenv("DB_USER_NAME")
	PASSWORD = os.Getenv("DB_PASSWORD")
	DATABASE = os.Getenv("DB_NAME")
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	driverName := "postgres"
	connectStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	db, err := sql.Open(driverName, connectStr)
	if err != nil {
		return fmt.Errorf("DB Open: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	if err := user.CreateUserTable(db); err != nil {
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

	err = http.ListenAndServe(":80", mux)
	if err != nil {
		return err
	}
	return nil
}
