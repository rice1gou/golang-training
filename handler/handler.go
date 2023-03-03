package handler

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "top page\n")
}
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	// サインイン処理
	fmt.Fprint(w, "signin\n")
}
func SignoutHandler(w http.ResponseWriter, r *http.Request) {
	// サインアウト処理
	fmt.Fprint(w, "signout\n")
}
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	// サインアップ処理
	fmt.Fprint(w, "signup\n")
}
func FetchUsersHandler(w http.ResponseWriter, r *http.Request) {
	// ユーザー一覧取得
	fmt.Fprint(w, "get users list\n")
}
func FetchUserHandler(w http.ResponseWriter, r *http.Request) {
	// ユーザー詳細取得
	fmt.Fprint(w, "get user details\n")
}
func SaveUserHandler(w http.ResponseWriter, r *http.Request) {
	// ユーザー登録
	fmt.Fprint(w, "register user\n")
}
func ModifyUserHandler(w http.ResponseWriter, r *http.Request) {
	// ユーザー編集
	fmt.Fprint(w, "modify user\n")
}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "delete user\n")
}
