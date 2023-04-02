package main

import (
	"net/http"
)

func main() {
	// Create a multiplexer
	mux := http.NewServeMux()

	// serve files out of the public directory
	files := http.FileServer(http.Dir("/public"))

	mux.Handle("/static/", http.StripPrefix("/static/", files))
	/* for any URL that starts with /static/ strip that prefix
	and look in the public directory for the file
	e.g. /static/css/main.css will look for /public/css/main.css */

	// redirects the root URL to a handler function
	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {

	threads, err := data.Threads()
	if err != nil {
		return
	}

	_, err = session(w, r)

	if err != nil {
		generateHTML(w, threads, "layout", "public.navbar", "index")
	} else {
		generateHTML(w, threads, "layout", "private.navbar", "index")
	}
}
