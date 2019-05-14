package main

import (
	"net/http"
)

func login(writer http.ResponseWriter, request *http.Request) {
	template := parseTemplateFiles("login.layout", "public.navbar", "login")
	template.Execute(writer, nil)
}

func signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}
