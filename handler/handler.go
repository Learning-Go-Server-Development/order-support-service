package handler

import "net/http"

type Handler interface {
	GetCustomer(w http.ResponseWriter, r *http.Request)
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetOrders(w http.ResponseWriter, r *http.Request)
	GetAddresses(w http.ResponseWriter, r *http.Request)
}
