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
	router := mux.NewRouter()
	//wiring
	//ch := CustomerHandlers{services.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{services.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//routes
	//mux.HandleFunc("/greet", greet).Methods(http.MethodGet)

	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomerById).Methods(http.MethodGet)
	//mux.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
