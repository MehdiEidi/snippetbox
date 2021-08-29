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

	log.Println("Starting server on :4000")
	log.Fatal(http.ListenAndServe(":4000", mux)) // starting a web server, listening on :4000(TCP network address)
}
