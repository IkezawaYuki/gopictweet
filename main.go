package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()



	mux.HandleFunc("/", index)

	mux.HandleFunc("/login", login)

	server := &http.Server{
		Addr:    config.Address,
		Handler: mux,
	}


	server.ListenAndServe()
}
