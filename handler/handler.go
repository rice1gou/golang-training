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
}
func SignoutHandler(w http.ResponseWriter, r *http.Request) {
	// サインアウト処理
	return
}
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	// サインアップ処理
	return
}
func FetchUsersHandler(w http.ResponseWriter, r *http.Request) {
	// ユーザー一覧取得
	return
}
func FetchUserHandler(w http.ResponseWriter, r *http.Request) {
	// ユーザー詳細取得
	return
}
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	// ユーザー登録
	return
}
func ModifyUserHandler(w http.ResponseWriter, r *http.Request) {
	// ユーザー編集
	return
}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	return
}
