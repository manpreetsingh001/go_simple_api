package main

import (
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/",version)
    //api(version)
    log.Fatal(http.ListenAndServe(":8080", nil))
}





