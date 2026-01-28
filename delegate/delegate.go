package delegate

type Delegate interface {
	GetCustomer(phone string) *Customer
	GetAddresses(cid int64) *[]Address
	GetOrders(cid int64) *[]Order
	GetProducts() *[]Item
}

type Service struct {
}

func (s *Service) Get() Delegate {
	return s
}
