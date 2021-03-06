package service

import (
	Dto "github.linwenqiang.com/GinStudy/Model/Dto"
	Dao "github.linwenqiang.com/GinStudy/Reponsetory/ORM"
	helper "github.linwenqiang.com/GinStudy/Tool"
)

type OrderService struct {
}

func (os *OrderService) CreateOrder(model *Dto.Order) int64 {
	dao := Dao.NewOrderDao()
	reflctrows, err := dao.CreateOrder(model)
	helper.PrintInfoError(err)
	return reflctrows
}
func (os *OrderService) UpdateOrder(model *Dto.Order) int64 {
	dao := Dao.NewOrderDao()
	reflctrows, err := dao.UpdateOrder(model)
	helper.PrintInfoError(err)
	return reflctrows
}

func (os *OrderService) QueryOrder(pageSearch Dto.GetOrderPageSearch) []Dto.GetOrderResult {
	dao := Dao.NewOrderDao()
	result, err := dao.QueryOrder(pageSearch)
	helper.PrintInfoError(err)
	return result
}

func (os *OrderService) QueryOrderDeatil(pageSearch Dto.GetOrderDetailPageSearch) []Dto.GetOrderDetailResult {
	dao := Dao.NewOrderDao()
	result, err := dao.QueryOrderDeatil(pageSearch)
	helper.PrintInfoError(err)
	return result
}
