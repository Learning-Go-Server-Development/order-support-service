package main

import (
	"fmt"
	"net/http"
)

func main() {

	port := "3001"
	msg := "Server starting on port "
	fmt.Println(msg + port)
	http.ListenAndServe(":3001", nil)

}

// go mod init github.com/Learning-Go-Server-Development/order-support-service
