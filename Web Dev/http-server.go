package main

import (
	"fmt"
	"net/http"
)

/*
	Bare minimum things an HTTP Server can do (demonstrated here):
	- process dynamic requests
	- serve static assets
	- accept connections
*/

func main() {

	/*
		Processing dynamic reuqests:
		http.Request lets you read GET/POST parameters
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!")
	})

	/*
		Serving static files:
		We have set up a file server with a path containing our static files
	*/
	fs := http.FileServer(http.Dir("static/"))

	// point a URL path to all of our files
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//accept connections like before
	http.ListenAndServe(":80", nil)
}
