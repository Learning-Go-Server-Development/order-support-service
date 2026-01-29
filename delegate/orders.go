package delegate

func (s *Service) GetOrders(cid int64) *[]Order {
	rtnList := []Order{}
	var olist []Order
	var o1 Order
	o1.ID = 12345555
	o1.CID = 12345
	var ilist []int64
	ilist = append(ilist, 12345, 258444)
	o1.Items = &ilist
	o1.OrderNumber = "OD-1255878"

	var o2 Order
	o2.ID = 2557788888
	o2.CID = 22345
	var ilist2 []int64
	ilist2 = append(ilist2, 12345, 66558)
	o2.Items = &ilist2
	o2.OrderNumber = "OD-1255879"

	var o3 Order
	o3.ID = 36654444
	o3.CID = 32345
	var ilist3 []int64
	ilist3 = append(ilist3, 12345, 4561122)
	o3.Items = &ilist3
	o3.OrderNumber = "OD-1255880"

	olist = append(olist, o1, o2, o3)

	for _, o := range olist {
		if o.CID == cid {
			rtnList = append(rtnList, o)
		}
	}

	return &rtnList
}
