package main

//订单记录
type OrderMainDAO interface {
	//保存
	SaveOrderMain()
	//DeleteOrderMain()
	//SearchOrderMain()
}

//订单详情
type OrderDetailDAO interface {
	//保存
	SaveOrderDetail()
}

type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}