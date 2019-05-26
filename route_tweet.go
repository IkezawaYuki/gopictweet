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
		generateHTML(w, nil, "layout", "private.navbar", "new.tweet")
	}
}

func editTweet(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	uuid := vals.Get("id")
	tweet, err := data.TweetByUuid(uuid)
	if err != nil {
		fmt.Println("error is occured")
	} else {
		_, err = session(w, r)
		if err != nil {
			http.Redirect(w, r, "/login", 302)
		} else {
			generateHTML(w, &tweet, "layout", "private.navbar", "edit.tweet")
		}
	}
}

func deleteTweet(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	uuid := vals.Get("id")
	tweet, err := data.TweetByUuid(uuid)
	if err != nil {
		fmt.Println("error is occured")
	} else {
		if err := tweet.Delete(tweet.Uuid); err != nil {
			fmt.Println("cannot delete tweet")
			panic(err)
		}
		http.Redirect(w, r, "/", 302)
	}
}

func updateTweet(w http.ResponseWriter, r *http.Request) {
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
		uuid := r.PostFormValue("uuid")
		text := r.PostFormValue("text")
		image := r.PostFormValue("image")
		if err := user.ModifyTweet(uuid, text, image); err != nil {
			fmt.Println("cannot uodate tweet")
			panic(err)
		}
		http.Redirect(w, r, "/", 302)
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

func createComment(w http.ResponseWriter, r *http.Request){
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}else{
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		user, err := sess.User()
		if err != nil {
			panic(err)
		}
		uuid := r.PostFormValue("uuid")
		text := r.PostFormValue("text")

		tweet, err := data.TweetByUuid(uuid)
		if err != nil {
			panic(err)
		}
		_, err = user.CreateComment(tweet, text)
		if err != nil {
			panic(err)
		}

		url := fmt.Sprint("/tweet/read?id=", uuid)
		http.Redirect(w, r, url, 302)
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
