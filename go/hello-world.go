package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    http.HandleFunc("/marco", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "POLO")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
