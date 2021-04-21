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
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w,"hostname: %s", host)
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
