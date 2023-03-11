package main

import (
	"database/sql"
	"log"
	"os"

	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/rice1gou/golang-training/cmd/identity/internal/handler"
	"github.com/rice1gou/golang-training/models/identity/user"
	"github.com/rice1gou/golang-training/pkg/router"
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

	ur := user.NewUserRepository(db)

	if err := ur.CreateUserTable(); err != nil {
		return err
	}

	mux := router.NewRouter()
	mux.Add(http.MethodGet, "/", handler.IndexHandler(ur))
	mux.Add(http.MethodPost, "/signin", handler.SigninHandler(ur))
	mux.Add(http.MethodGet, "/signout", handler.SignoutHandler(ur))
	mux.Add(http.MethodGet, "/signup", handler.SignupHandler(ur))
	mux.Add(http.MethodGet, "/user", handler.FetchUsersHandler(ur))
	mux.Add(http.MethodPost, "/user", handler.SaveUserHandler(ur))
	mux.Add(http.MethodGet, "/user/([^/]+)", handler.FetchUserDetailsHandler(ur))
	mux.Add(http.MethodPost, "/user/([^/]+)", handler.ModifyUserHandler(ur))
	mux.Add(http.MethodDelete, "/user/([^/]+)", handler.DeleteUserHandler(ur))

	err = http.ListenAndServe(":80", mux)
	if err != nil {
		return err
	}
	return nil
}
