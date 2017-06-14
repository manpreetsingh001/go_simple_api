package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func version(w http.ResponseWriter, r *http.Request) {
	to := Todos{
		available{Name: "quake", Lang: "golangbuild", Ver: []int{1, 1}},
	}
	log.Printf("Request for: %s", r.URL.Path)
	fmt.Println(to)
	if err := json.NewEncoder(w).Encode(to); err != nil {
		panic(err)
	}

}
