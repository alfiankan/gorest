package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customer", getAllCustomer)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}
