package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define a new command-line flag with the name 'addr', a default value of ":4000"
	// and some short help text explaining what the flag controls. The value of the
	// flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this *before* you use the addr variable
	// otherwise it will always contain the default value of ":4000". If any errors are
	// encountered during parsing the application will be terminated.
	flag.Parse()

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

	// The value returned from the flag.String() function is a pointer to the flag
	// value, not the value itself. So we need to dereference the pointer (i.e.
	// prefix it with the * symbol) before using it.
	log.Println("Starting server on", *addr)
	log.Fatal(http.ListenAndServe(*addr, mux)) // starting a web server, listening on *addr(TCP network address)
}
