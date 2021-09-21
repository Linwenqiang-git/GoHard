package service

import (
	Dto "github.linwenqiang.com/GinStudy/Model/Dto"
	Dao "github.linwenqiang.com/GinStudy/Reponsetory/ORM"
	helper "github.linwenqiang.com/GinStudy/Tool"
)

type OrderService struct {
}

func (os *OrderService) CreateOrder(model Dto.OrderDto) int64 {
	dao := Dao.NewOrderDao()
	reflctrows, err := dao.CreateOrder(model)
	helper.PrintInfoError(err)
	return reflctrows
}

func (os *OrderService) QueryOrder(pageSearch Dto.GetOrderPageSearch) []Dto.GetOrderResult {
	dao := Dao.NewOrderDao()
	result, err := dao.QueryOrder(pageSearch)
	helper.PrintInfoError(err)
	return result
}
