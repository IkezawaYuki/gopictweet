package main

import "net/http"

func newTweet(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, r, "layout", "public.navbar", "new.tweet")
	}
}
