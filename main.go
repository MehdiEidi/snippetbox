package main

import (
	"log"
	"net/http"
)

// homeHandler is the handler for "/" URL pattern
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// For the reason that "/" in serve mux acts as catch-all(subtree pattern). So, we make
	// the homeHandler return 404 not found response to all the patterns starting with "/" (not registered ones)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from snippet box"))
}

// showSnippetHandler is the handler for "/snippet" URL pattern
func showSnippetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet"))
}

// createSnippetHandler is the handler for "/snippet/creat" URl pattern
func createSnippetHandler(w http.ResponseWriter, r *http.Request) {
	// This handler should only function when the request method is POST
	// otherwise should return 405 header: method not allowed.
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST") // Adding "Allow: POST" header to response header map. To show user which method is allowed
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	w.Write([]byte("create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()                               // initialize a new servemux
	mux.HandleFunc("/", homeHandler)                        // register homeHandler as the handler for "/" URL pattern
	mux.HandleFunc("/snippet", showSnippetHandler)          // register showSnippetHandler as the handler for "/snippet" URL pattern
	mux.HandleFunc("/snippet/create", createSnippetHandler) // register createSnippetHandler as the handler for "/snippet/create" URL pattern

	log.Println("Starting server on :4000")
	log.Fatal(http.ListenAndServe(":4000", mux)) // starting a web server, listening on :4000(TCP network address)
}
