package main

import (
	"errors"
	"fmt"
	"gopictweet/data"
	"html/template"
	"log"
	"net/http"
)

var logger *log.Logger

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	fmt.Println(templates)
	templates.ExecuteTemplate(writer, "layout", data)
}

func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(template.ParseFiles(files...))
	return
}

func session(w http.ResponseWriter, r *http.Request) (ses data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		ses = data.Session{Uuid: cookie.Value}
		if ok, _ := ses.Check(); !ok {
			err = errors.New("invalid error")
		}
	}
	return
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR")
	logger.Println(args...)
}
