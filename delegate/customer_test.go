package delegate_test

import (
	"testing"

	del "github.com/Learning-Go-Server-Development/order-support-service/delegate"
)

func TestService_GetCustomer(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		phone string
		want  int64
	}{
		// TODO: Add test cases.
		{
			name:  "test1",
			phone: "954-555-7858",
			want:  22345,
		},
		{
			name:  "test2",
			phone: "154-555-7878",
			want:  12345,
		},
		{
			name:  "test3",
			phone: "678-656-7878",
			want:  32345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var s del.Service
			d := s.Get()
			got := d.GetCustomer(tt.phone)
			// TODO: update the condition below to compare got with tt.want.
			if got.ID != tt.want {
				t.Errorf("GetCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}
