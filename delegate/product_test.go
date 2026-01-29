package delegate_test

import (
	"testing"

	"github.com/Learning-Go-Server-Development/order-support-service/delegate"
)

func TestService_GetProducts(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var s delegate.Service
			d := s.Get()
			got := d.GetProducts()
			// TODO: update the condition below to compare got with tt.want.
			if len(*got) != 4 {
				t.Errorf("GetProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}
