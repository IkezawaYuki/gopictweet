package main

import (
	"net/http"

	"gopictweet/data"
)

func index(writer http.ResponseWriter, response *http.Request) {
	tweets, err := data.Tweets()
	if err != nil {
		generateHTML(writer, tweets, "layout", "public.navbar", "index")
	}
}

func err(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, vals.Get("Msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(w, vals.Get("Msg"), "layout", "private.navbar", "error")
	}
}
