package service

//工厂类，负责管理service层的创建
func NewOrdeService() *OrderService {
	return &OrderService{}
}
