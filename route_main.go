package main

import (
	"fmt"
	"net/http"

	"./data"
)

func index(writer http.ResponseWriter, response *http.Request) {
	tweets, err := data.Tweets()
	if err != nil {
		fmt.Println("not err!")
		generateHTML(writer, tweets, "layout", "public.navbar", "index")
	}
}
