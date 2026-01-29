package delegate

func (s *Service) GetProducts() *[]Item {
	var i1 Item
	i1.ID = 12345
	i1.SKU = "123444"
	i1.Description = "System76 Laptop"
	i1.Price = 1924.00

	var i2 Item
	i2.ID = 258444
	i2.SKU = "2558444"
	i2.Description = "System76 Monitor"
	i2.Price = 798.00

	var i3 Item
	i3.ID = 66558
	i3.SKU = "6655444777"
	i3.Description = "System76 Mouse"
	i3.Price = 78.00

	var i4 Item
	i4.ID = 474748
	i4.SKU = "4561122"
	i4.Description = "System76 Hard Drive 1TB"
	i4.Price = 285.00

	var ilist []Item
	ilist = append(ilist, i1, i2, i3, i4)

	return &ilist
}
