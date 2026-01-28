package delegate

func (s *Service) GetAddresses(cid int64) *[]Address {
	var rtn []Address

	var a1 Address
	a1.ID = 111
	a1.CID = 12345
	a1.Street = "121 Peachtree St"
	a1.City = "Atlanta"
	a1.State = "GA"
	a1.ZipCode = "23547"

	var a2 Address
	a2.ID = 222
	a2.CID = 22345
	a2.Street = "1214 Franklin St"
	a2.City = "Atlanta"
	a2.State = "GA"
	a2.ZipCode = "23544"

	var a3 Address
	a3.ID = 333
	a3.CID = 32345
	a3.Street = "2555 Lee St"
	a3.City = "Atlanta"
	a3.State = "GA"
	a3.ZipCode = "35565"

	var alist []Address
	alist = append(alist, a1, a2, a3)

	for _, a := range alist {
		if a.CID == cid {
			rtn = append(rtn, a)
		}
	}

	return &rtn
}
