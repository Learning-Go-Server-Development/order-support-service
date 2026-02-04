package delegate_test

import (
	"testing"

	"github.com/Learning-Go-Server-Development/order-support-service/delegate"
)

func TestService_GetProducts(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		sku  string
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			sku:  "2558444",
			want: 258444,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var s delegate.Service
			d := s.Get()
			got := d.GetProduct(tt.sku)
			// TODO: update the condition below to compare got with tt.want.
			if got.ID != tt.want {
				t.Errorf("GetProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}
