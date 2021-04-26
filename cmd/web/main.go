package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// This essentially defines a new command-line flag with the name addr, a
	// default value of ":4000" and some short help text explaining what the
	// flag controls. The value of the flag will be stored in the addr variable at
	// runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// flag.Parse() parse the command-line arguments update default
	// values of the flag with provided value.
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Create a file server which serves files out of the "./ui/static" directory
	// Note that the path given to the http.Dir function is relative to project's
	// root directory
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// The value returned from the flag.String() is pointer
	// So addr variable is pointer to a string value. We need to
	// dereference the pointer.
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
