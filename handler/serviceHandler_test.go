package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mux "github.com/GolangToolKits/grrt"
	del "github.com/Learning-Go-Server-Development/order-support-service/delegate"
	han "github.com/Learning-Go-Server-Development/order-support-service/handler"
)

func TestServiceHandler_GetCustomer(t *testing.T) {
	var sh han.ServiceHandler
	var s del.Service
	d := s.Get()
	sh.Deligate = d
	h := sh.New()

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		phoneNumber string
		w           http.ResponseWriter
		r           *http.Request
	}{
		// TODO: Add test cases.
		{
			name:        "test 1",
			phoneNumber: "954-555-7858",
		},
		{
			name:        "test 2",
			phoneNumber: "",
		},
		{
			name:        "test 3",
			phoneNumber: "954-555-1111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := http.NewRequest("GET", "/getStuff", nil)
			vars := map[string]string{
				"phoneNumber": tt.phoneNumber,
			}
			r = mux.SetURLVars(r, vars)
			w := httptest.NewRecorder()
			h.GetCustomer(w, r)
			if tt.name == "test 1" && w.Code != 200 {
				t.Fail()
			} else if tt.name == "test 2" && w.Code != 400 {
				t.Fail()
			} else if tt.name == "test 3" && w.Code != 200 {
				t.Fail()
			}
		})
	}
}

func TestServiceHandler_GetProducts(t *testing.T) {
	var sh han.ServiceHandler
	var s del.Service
	d := s.Get()
	sh.Deligate = d
	h := sh.New()

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			r, _ := http.NewRequest("GET", "/getStuff", nil)
			w := httptest.NewRecorder()
			h.GetProducts(w, r)
			if tt.name == "test 1" && w.Code != 200 {
				t.Fail()
			}
		})
	}
}
