package handler

import (
	"database/sql"
	"fmt"
	"net/http"
)

func IndexHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "top page\n")
	}
}

func SigninHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	// サインイン処理
	fmt.Fprint(w, "signin\n")
	}
}

func SignoutHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
	// サインアウト処理
	fmt.Fprint(w, "signout\n")
	}
}

func SignupHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
	// サインアップ処理
	fmt.Fprint(w, "signup\n")
	}
}
func FetchUsersHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
	// ユーザー一覧取得
	fmt.Fprint(w, "get users list\n")

	}
}

func FetchUserHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
	// ユーザー詳細取得
	fmt.Fprint(w, "get user details\n")
	}
}

func SaveUserHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
	// ユーザー登録
	fmt.Fprint(w, "register user\n")
	}
}

func ModifyUserHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
	// ユーザー編集
	fmt.Fprint(w, "modify user\n")
	}
}

func DeleteUserHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "delete user\n")
	}
}
