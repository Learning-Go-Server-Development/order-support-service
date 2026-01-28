package delegate

func (s *Service) GetCustomer(phone string) *Customer {
	var rtn Customer
	var c1 Customer
	c1.ID = 12345
	c1.FirstName = "Bob"
	c1.LastName = "Roberts"
	c1.PhoneNumber = "154-555-7878"

	var c2 Customer
	c2.ID = 22345
	c2.FirstName = "Jim"
	c2.LastName = "Snaders"
	c2.PhoneNumber = "954-555-7858"

	var c3 Customer
	c3.ID = 32345
	c3.FirstName = "William"
	c3.LastName = "Getts"
	c3.PhoneNumber = "678-656-7878"

	var clist []Customer
	clist = append(clist, c1)
	clist = append(clist, c2)
	clist = append(clist, c3)

	for _, c := range clist {
		if c.PhoneNumber == phone {
			rtn = c
		}
	}
	return &rtn
}
