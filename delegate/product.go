package delegate

func (s *Service) GetProducts() *[]Item {
	var i1 Item
	i1.ID = 12345
	i1.SKU = "123444"
	i1.Description = "System76 Laptop"
	i1.Price = 1924.00

	return nil
}
