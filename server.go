package main

import (
	"fmt"
	"net/http"

	mux "github.com/GolangToolKits/grrt"
	del "github.com/Learning-Go-Server-Development/order-support-service/delegate"
	han "github.com/Learning-Go-Server-Development/order-support-service/handler"
)

func main() {
	// use dependency injection here
	var sh han.ServiceHandler
	var s del.Service

	d := s.Get()
	sh.Deligate = d
	h := sh.New()
	router := mux.NewRouter()

	router.HandleFunc("/rs/customer/get/{phoneNumber}", h.GetCustomer).Methods("GET")
	router.HandleFunc("/rs/product/get/{sku}", h.GetProduct).Methods("GET")
	router.HandleFunc("/rs/orders/get/{cid}", h.GetOrders).Methods("GET")
	router.HandleFunc("/rs/addresses/get/{cid}", h.GetAddresses).Methods("GET")

	port := "3001"
	msg := "Server starting on port "
	fmt.Println(msg + port)
	http.ListenAndServe(":3001", router)

}

// go mod init github.com/Learning-Go-Server-Development/order-support-service
