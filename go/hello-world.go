package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {


    http.HandleFunc("/MARCO", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "POLO!")
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf("hostname:", name)
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
