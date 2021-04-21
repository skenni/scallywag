package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {


    http.HandleFunc("/marco", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "\n\nPOLO!")
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w,"\n\nhostname: %s", host)
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
