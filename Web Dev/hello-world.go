package main

import (
	"fmt"
	"net/http"
)

func main() {

	/*
		signature for request handlers is: func (w http.ResponseWriter, r *http.Request)

		http.ResponseWriter is where we write our responses to
		http.Request contians information about our request (url, header field, etc.)

	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello you've requested: %s\n", r.URL.Path)
	})
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	/*
		We need to specify a port to pass connections on to the request
		handler
	*/
	http.ListenAndServe(":80", nil)
}
