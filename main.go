package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
    // Check if the current request URL path exactly matches "/".
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.Header().Set("Allow", "POST")
        http.Error(w, "Method Not Allowed", 405)
        return
    }

    w.Write([]byte("Create a new snipet..."))
}

func main() {
    // Resister the two new handler functions and corresponding URL pattern with
    // servemux.
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet", showSnippet)
    mux.HandleFunc("/snippet/create", createSnippet)

    // Use the  http.ListenAndServe() function to start a new web server.
    // We pass two parameters: The TCP network address to listen on (in
    // this case ":4000" and the servemux we just created.
    log.Println("Starting sever on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
