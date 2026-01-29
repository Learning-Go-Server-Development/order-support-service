package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	del "github.com/Learning-Go-Server-Development/order-support-service/delegate"
)

type Handler interface {
	GetCustomer(w http.ResponseWriter, r *http.Request)
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetOrders(w http.ResponseWriter, r *http.Request)
	GetAddresses(w http.ResponseWriter, r *http.Request)
}

type ServiceHandler struct {
	Deligate del.Delegate
}

func (h *ServiceHandler) New() Handler {
	return h
}

// CheckContent CheckContent
func (h *ServiceHandler) CheckContent(r *http.Request) bool {
	var rtn bool
	cType := r.Header.Get("Content-Type")
	if cType == "application/json" {
		// http.Error(w, "json required", http.StatusUnsupportedMediaType)
		rtn = true
	}
	return rtn
}

// SetContentType SetContentType
func (h *ServiceHandler) SetContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// ProcessBody ProcessBody
func (h *ServiceHandler) ProcessBody(r *http.Request, obj any) (bool, error) {
	var suc bool
	var err error
	//fmt.Println("r.Body: ", r.Body)
	//log.Debug("r.Body: ", r.Body)
	if r.Body != nil {
		decoder := json.NewDecoder(r.Body)
		//fmt.Println("decoder: ", decoder)
		err = decoder.Decode(obj)
		//fmt.Println("decoder: ", decoder)
		if err != nil {
			//log.Println("Decode Error: ", err.Error())
			log.Println("Decode Error: ", err.Error())
		} else {
			suc = true
		}
	} else {
		err = errors.New("Bad Body")
	}
	return suc, err
}
