package delegate_test

import (
	"testing"

	"github.com/Learning-Go-Server-Development/order-support-service/delegate"
)

func TestService_GetOrders(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		cid   int64
		want  int64
		want2 int
	}{
		// TODO: Add test cases.
		{
			name:  "test 1",
			cid:   22345,
			want:  2557788888,
			want2: 2,
		},
		{
			name:  "test 2",
			cid:   12345,
			want:  12345555,
			want2: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var s delegate.Service
			d := s.Get()
			got := d.GetOrders(tt.cid)
			// TODO: update the condition below to compare got with tt.want.
			// i:= (*got)[0].Items
			if (*got)[0].ID != tt.want && len(*(*got)[0].Products) != tt.want2 {
				t.Errorf("GetOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
