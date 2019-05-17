package main

import (
	"fmt"
	"gopictweet/data"
	"net/http"
)

func newTweet(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, r, "layout", "public.navbar", "new.tweet")
	}
}

func createTweet(w http.ResponseWriter, r *http.Request) {
	ses, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			fmt.Println("cannnot parse")
			panic(err)
		}
		user, err := ses.User()
		if err != nil {
			panic(err)
		}
		text := r.PostFormValue("text")
		image := r.PostFormValue("image")
		if _, err := user.CreateTweet(text, image); err != nil {
			fmt.Println("cannot create tweet")
			panic(err)
		}
		http.Redirect(w, r, "/", 302)
	}
}

func readTweet(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	uuid := vals.Get("id")
	tweet, err := data.TweetByUuid(uuid)
	if err != nil {
		fmt.Println("error is occured")
	} else {
		_, err = session(w, r)
		if err != nil {
			generateHTML(w, &tweet, "layout", "public.nuvbar", "public.tweet")
		} else {
			generateHTML(w, &tweet, "layout", "private.navbar", "private.tweet")
		}
	}
}
