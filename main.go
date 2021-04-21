package main

import (
    "log"
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from Snippetbox"))
}

func main() {
    // Initialize a new servemux, then
    // register the home function as the handler for the "/" URL pattern.
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)

    // Use the  http.ListenAndServe() function to start a new web server.
    // We pass two parameters: The TCP network address to listen on (in
    // this case ":4000" and the servemux we just created.
    log.Println("Starting sever on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
