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
