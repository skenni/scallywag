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
        fmt.Fprintf(w, "goodbye")
    })

    http.HandleFunc("/MARCO", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "POLO")
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
