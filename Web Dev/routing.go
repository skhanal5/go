package main

/*
	gorilla/mux is a pacakge that adds functionality to Go's HTTP router
	while adhering to its default request handler signature
*/

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	/*
		Creates a new request router

		Receives all HTTP connections and pass it onto the
		request handlers we register on it
	*/
	r := mux.NewRouter()

	/*
		Registering a route handler

		The following route has two dynamic segments: {title} and {page}
		We can extract those segments easily by referencing the slugs

	*/
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, f *http.Request) {
		vars := mux.Vars(f)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "This is the title: %s and this is the page: %s", title, page)
	})

	/*
		Remember the second paramter was nil.
		Now it's r, because we want to use a different
		router in place of the default one
	*/
	http.ListenAndServe(":80", r)
}
