package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	noOfBytes, err := fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
	/* request.URL.Path
	prints out the path of the url after the port number
	and before the query string
	and shows it in the browser.
	[1:] is used to remove the first character of the string, (the initial slash) */

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(noOfBytes)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
