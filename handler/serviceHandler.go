package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	mux "github.com/GolangToolKits/grrt"
	del "github.com/Learning-Go-Server-Development/order-support-service/delegate"
)

func (h *ServiceHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	h.SetContentType(w)
	vars := mux.Vars(r)
	log.Println("vars: ", len(vars))
	if len(vars) == 1 {
		phone := vars["phoneNumber"]
		log.Println("phoneNumber:", phone)
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
	h.SetContentType(w)
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
	h.SetContentType(w)
	vars := mux.Vars(r)
	log.Println("vars: ", len(vars))
	if len(vars) == 1 {
		cidStr := vars["cid"]
		if cidStr == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			var orders *[]del.Order
			cid, ciderr := strconv.ParseInt(cidStr, 10, 64)
			if ciderr == nil {
				//implementation of delegate is hidden here
				orders = h.Deligate.GetOrders(cid)
				if orders != nil {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusBadRequest)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}

			resJSON, _ := json.Marshal(orders)
			fmt.Fprint(w, string(resJSON))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *ServiceHandler) GetAddresses(w http.ResponseWriter, r *http.Request) {
	h.SetContentType(w)
	vars := mux.Vars(r)
	log.Println("vars: ", len(vars))
	if len(vars) == 1 {
		cidStr := vars["cid"]
		if cidStr == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			var addresses *[]del.Address
			cid, ciderr := strconv.ParseInt(cidStr, 10, 64)
			if ciderr == nil {
				//implementation of delegate is hidden here
				addresses = h.Deligate.GetAddresses(cid)
				if addresses != nil {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusBadRequest)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
			resJSON, _ := json.Marshal(addresses)
			fmt.Fprint(w, string(resJSON))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
