package main

import (
	"net/http"

	"gopictweet/data"
)

func index(w http.ResponseWriter, r *http.Request) {
	tweets, err := data.Tweets()
	if err != nil{
		error_message(w, r, "cannnot get tweets")
	}else{
		_, err = session(w, r)
		if err != nil{
			generateHTML(w, tweets, "layout", "public.navbar", "index")
		}else{
			generateHTML(w, tweets, "layout", "private.navbar", "index")
		}
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
