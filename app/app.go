package app

import (
	"log"
	"net/http"

	"github.com/alfiankan/gorest/domain"
	"github.com/alfiankan/gorest/services"
	"github.com/gorilla/mux"
)

func Start() {
	//mux := http.NewServeMux()
	mux := mux.NewRouter()
	//wiring
	ch := CustomerHandlers{services.NewCustomerService(domain.NewCustomerRepositoryStub())}
	//routes
	mux.HandleFunc("/greet", greet).Methods(http.MethodGet)

	mux.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	mux.HandleFunc("/customer/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	mux.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}
