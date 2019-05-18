package main

import (
	"net/http"
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

	mux.HandleFunc("/tweet/new", newTweet)
	mux.HandleFunc("/tweet/read", readTweet)
	mux.HandleFunc("/tweet/create", createTweet)
	mux.HandleFunc("/tweet/new", newTweet)


	server := &http.Server{
		Addr:    config.Address,
		Handler: mux,
	}


	server.ListenAndServe()
}
