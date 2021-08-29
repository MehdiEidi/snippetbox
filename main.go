package main

import (
	"log"
	"net/http"
)

// homeHandler is the handler for "/" URL pattern
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippet box"))
}

func main() {
	mux := http.NewServeMux()        // initialize a new servemux
	mux.HandleFunc("/", homeHandler) // register homeHandler as the handler for "/" URL pattern

	log.Println("Starting server on :4000")
	log.Fatal(http.ListenAndServe(":4000", mux)) // starting a web server, listening on :4000(TCP network address)
}
