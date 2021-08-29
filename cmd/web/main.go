package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()                               // initialize a new servemux
	mux.HandleFunc("/", homeHandler)                        // register homeHandler as the handler for "/" URL pattern
	mux.HandleFunc("/snippet", showSnippetHandler)          // register showSnippetHandler as the handler for "/snippet" URL pattern
	mux.HandleFunc("/snippet/create", createSnippetHandler) // register createSnippetHandler as the handler for "/snippet/create" URL pattern

	// Create a file server which serves files out of the "./ui/static" directory
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on :4000")
	log.Fatal(http.ListenAndServe(":4000", mux)) // starting a web server, listening on :4000(TCP network address)
}
