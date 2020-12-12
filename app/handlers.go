package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/alfiankan/gorest/services"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" cml:"name"`
	City    string
	Zipcode string
}

type CustomerHandlers struct {
	services services.CustomerService
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func (ch *CustomerHandlers) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"Alfiankan", "Solo", "577871"},
	// 	{"Supardi", "Bandung", "577874"},
	// }
	customers,_ := ch.services.GetAllCustomer()
	//tambah info di header kalau response berupa json app
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)

	}

}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request accepted")
}
