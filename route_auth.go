package main

import (
	"fmt"
	"gopictweet/data"
	"net/http"
)

func login(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "login.layout", "public.navbar", "login")
}

func signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

func signupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	user := data.User{
		Nickname: r.PostFormValue("nickname"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	if err = user.Create(); err != nil {
		panic(err)
	}
	http.Redirect(w, r, "login", 302)
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		// todo ログインできません的なメッセージを出す。
		//danger("connot find user")
	}
	fmt.Println("danger is throwgn")
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			panic(err)
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)

	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		session := data.Session{Uuid: cookie.Value}
		err = session.DeleteByUUID()
	}
	http.Redirect(w, r, "/", 302)
}
