package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/rice1gou/golang-training/pkg/router"
	"github.com/rice1gou/golang-training/pkg/user"
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
		users, err := user.FetchUsers(db)
		if err != nil {
			http.Error(w, "users not found", http.StatusInternalServerError)
		}
		for _, u := range users {
			fmt.Fprintf(w, "userid:%s, username:%s\n", u.UserId, u.UserName)
		}
	}
}

func FetchUserDetailsHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		// ユーザー詳細取得
		fmt.Fprint(w, "get user details\n")
		oid := router.PathParam(r, 0)
		u, err := user.FetchUserDetails(db, oid)
		if err != nil {
			http.Error(w, "User Not Exists", http.StatusBadRequest)
		}
		fmt.Fprintf(w, "useroid: %s, userid: %s, username: %s", oid, u.UserId, u.UserName)
	}
}

func SaveUserHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		// ユーザー登録
		fmt.Fprint(w, "register user\n")
		r.ParseForm()
		body := r.PostForm
		uid := body["userid"][0]
		uname := body["username"][0]
		pw := body["password"][0]

		u := user.NewUser(uid, uname, pw)
		if err := user.SaveUser(db, *u); err != nil {
			http.Error(w, "couldn't saved user", http.StatusInternalServerError)
		}
	}
}

func ModifyUserHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		// ユーザー編集
		fmt.Fprint(w, "modify user\n")
		r.ParseForm()
		body := r.PostForm
		var u user.User
		u.UserOid = router.PathParam(r, 0)
		u.UserId = body["userid"][0]
		u.UserName = body["username"][0]

		err := user.ModifyUser(db, &u)
		if err != nil {
			http.Error(w, "couldn't update user", http.StatusInternalServerError)
		}
	}
}

func DeleteUserHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "delete user\n")
		oid := router.PathParam(r, 0)
		err := user.DeleteUser(db, oid)
		if err != nil {
			http.Error(w, "couldn't delete user", http.StatusInternalServerError)
		}
	}
}
