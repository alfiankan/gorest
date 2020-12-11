package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	//mux := http.NewServeMux()
	mux := mux.NewRouter()
	mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers", getAllCustomer).Methods(http.MethodGet)
	mux.HandleFunc("/customer/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	mux.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}
