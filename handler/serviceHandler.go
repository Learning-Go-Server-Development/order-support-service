package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mux "github.com/GolangToolKits/grrt"
)

func (h *ServiceHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("vars: ", len(vars))
	if len(vars) == 1 {
		phone := vars["phoneNumber"]
		if phone == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			//implementation of delegate is hidden here
			customer := h.Deligate.GetCustomer(phone)
			if customer != nil {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
			resJSON, _ := json.Marshal(customer)
			fmt.Fprint(w, string(resJSON))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *ServiceHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := h.Deligate.GetProducts()
	if products != nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	resJSON, _ := json.Marshal(products)
	fmt.Fprint(w, string(resJSON))
}

func (h *ServiceHandler) GetOrders(w http.ResponseWriter, r *http.Request) {

}

func (h *ServiceHandler) GetAddresses(w http.ResponseWriter, r *http.Request) {

}
