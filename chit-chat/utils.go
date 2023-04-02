package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	// retrieve cookie from request
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		err = errors.New("Invalid session")
		return
	}

	sess = data.Session{Uuid: cookie.Value}
	if ok, _ := sess.Check(); !ok {
		err = errors.New("Invalid session")
	}
	return
}

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// func parseTemplateFiles(filenames ...string) (t *template.Template) {
// 	var files []string
// 	for _, file := range
// }
