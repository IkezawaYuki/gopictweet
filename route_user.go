package main

import (
	"fmt"
	"gopictweet/data"
	"net/http"
)

func mypage(w http.ResponseWriter, r *http.Request){
	vals := r.URL.Query()
	uuid := vals.Get("id")
	user, err := data.UserByUUID(uuid)
	if err != nil {
		fmt.Println("error is occured")
	} else {
		_, err = session(w, r)
		if err != nil {
			http.Redirect(w, r, "/login", 302)
		} else {
			generateHTML(w, &user, "layout", "private.navbar", "private.user")
		}
	}
}

func userpage(w http.ResponseWriter, r *http.Request){
	vals := r.URL.Query()
	uuid := vals.Get("id")
	user, err := data.UserByUUID(uuid)
	fmt.Println(user)
	if err != nil {
		fmt.Println("error is occured")
	} else {
		_, err = session(w, r)
		if err != nil {
			http.Redirect(w, r, "/login", 302)
		} else {
			generateHTML(w, &user, "layout", "public.navbar", "public.user")
		}
	}
}