package handler

import (
	"fmt"
	"net/http"

	"github.com/rice1gou/golang-training/models/identity/user"
	"github.com/rice1gou/golang-training/pkg/router"
)

func IndexHandler(ur user.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "top page\n")
	}
}

func SigninHandler(ur user.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// サインイン処理
		fmt.Fprint(w, "signin\n")
	}
}

func SignoutHandler(ur user.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// サインアウト処理
		fmt.Fprint(w, "signout\n")
	}
}

func SignupHandler(ur user.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// サインアップ処理
		fmt.Fprint(w, "signup\n")
	}
}

func FetchUsersHandler(ur user.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ユーザー一覧取得
		fmt.Fprint(w, "get users list\n")
		users, err := ur.FetchUsers()
		if err != nil {
			http.Error(w, "users not found", http.StatusInternalServerError)
		}
		for _, u := range users {
			fmt.Fprintf(w, "useroid:%s, userid:%s, username:%s\n", u.UserOid, u.UserId, u.UserName)
		}
	}
}

func FetchUserDetailsHandler(ur user.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ユーザー詳細取得
		fmt.Fprint(w, "get users details\n")
		oid := router.PathParam(r, 0)
		u, err := ur.FetchUserDetails(oid)
		if err != nil {
			http.Error(w, "User Not Exists", http.StatusBadRequest)
		}
		fmt.Fprintf(w, "useroid: %s, userid: %s, username: %s", oid, u.UserId, u.UserName)
	}
}

func SaveUserHandler(ur user.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ユーザー登録
		fmt.Fprint(w, "register user\n")
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Form request parse error", http.StatusBadRequest)
		}
		uid := r.Form.Get("userid")
		uname := r.Form.Get("username")
		pw := r.Form.Get("password")
		fmt.Fprintf(w, "userid: %s, username: %s, password: %s", uid, uname, pw)

		u := user.NewUser(uid, uname, pw)
		if err := ur.SaveUser(u); err != nil {
			http.Error(w, "couldn't saved user", http.StatusInternalServerError)
		}
	}
}

func ModifyUserHandler(ur user.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ユーザー編集
		fmt.Fprint(w, "modify user\n")
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Form request parse error", http.StatusBadRequest)
		}
		body := r.PostForm
		userOid := router.PathParam(r, 0)
		userId := body["userid"][0]
		userName := body["username"][0]

		err := ur.ModifyUser(userOid, userId, userName)
		if err != nil {
			http.Error(w, "couldn't update user", http.StatusInternalServerError)
		}
	}
}

func DeleteUserHandler(ur user.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "delete user\n")
		oid := router.PathParam(r, 0)
		err := ur.DeleteUser(oid)
		if err != nil {
			http.Error(w, "couldn't delete user", http.StatusInternalServerError)
		}
	}
}
