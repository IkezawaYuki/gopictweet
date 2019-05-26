package main

import (
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signout", logout)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/user/mypage", mypage)

	mux.HandleFunc("/tweet/new", newTweet)
	mux.HandleFunc("/tweet/read", readTweet)
	mux.HandleFunc("/tweet/create", createTweet)
	mux.HandleFunc("/tweet/edit", editTweet)
	mux.HandleFunc("/tweet/update", updateTweet)
	mux.HandleFunc("/tweet/delete", deleteTweet)

	mux.HandleFunc("/tweet/comment", createComment)

	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
