package main

import (
	"net/http"
)

// GET /login
// Show the login page
// func login(w http.ResponseWriter, r *http.Request){
// 	t := parseTemplateFiles("login.layout", "public.navbar", "login")
// 	t.Execute(w, nil)
// }

func authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, _ := data.UserByEmail(r.PostFormValue("email"))

	if user.Password != data.Encrypt(r.PostFormValue("password")) {
		http.Redirect(w, r, "/login", 302)
	}

	session := user.CreateSession()
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.Uuid,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", 302)
}
