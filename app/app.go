package main

import (
	"log"
	"net/http"
)

func Start() {

	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customer", getAllCustomer)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
