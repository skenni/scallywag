package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "WAZZUP, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "GOODBYE")
    })

    http.HandleFunc("/MARCO", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Polo!")
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
