package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// homeHandler is the handler for "/" URL pattern
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// For the reason that "/" in serve mux acts as catch-all(subtree pattern). So, we make
	// the homeHandler return 404 not found response to all the patterns starting with "/" (not registered ones)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Initialize a slice containing the paths to the two files. Note that the
	// home.page.tmpl file must be the *first* file in the slice.
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// Use the template.ParseFiles() function to read the template files into a
	// template set
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// We then use Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents any
	// dynamic data that we want to pass in, which for now we'll leave as nil.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// showSnippetHandler is the handler for "/snippet" URL pattern
func showSnippetHandler(w http.ResponseWriter, r *http.Request) {
	// Getting the id value from query string. Also converting it to int to see if id is actually int and later
	// check if it is a positive number
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display snippet with ID:%d", id)
}

// createSnippetHandler is the handler for "/snippet/creat" URl pattern
func createSnippetHandler(w http.ResponseWriter, r *http.Request) {
	// This handler should only function when the request method is POST
	// otherwise should return 405 header: method not allowed.
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST") // Adding "Allow: POST" header to response header map. To show user which method is allowed

		// This code below is equivalent to these two: w.WriteHeader(http.StatusMethodNotAllowed) then w.Write([]byte("Method not allowed"))
		http.Error(w, "Method not allowed", 405)
		return
	}

	w.Write([]byte("create a new snippet..."))
}
