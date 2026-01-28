package delegate_test

import (
	"testing"

	"github.com/Learning-Go-Server-Development/order-support-service/delegate"
)

func TestService_GetAddresses(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		cid  int64
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			cid:  22345,
			want: 222,
		},
		{
			name: "test2",
			cid:  32345,
			want: 333,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var s delegate.Service
			d := s.Get()
			got := d.GetAddresses(tt.cid)
			// TODO: update the condition below to compare got with tt.want.
			if (*got)[0].ID != tt.want {
				t.Errorf("GetAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
