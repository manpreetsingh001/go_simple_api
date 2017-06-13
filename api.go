package main

import (

	"net/http"
)

type path func(w http.ResponseWriter, r *http.Request)

func api(p path ) {
	http.HandleFunc("/",p)

}






