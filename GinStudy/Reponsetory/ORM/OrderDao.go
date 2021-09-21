package orm

import (
	"github.com/go-xorm/xorm"
	Dto "github.linwenqiang.com/GinStudy/Model/Dto"
)

//为orderService提供数据库访问层
type OrderDao struct {
	engine *xorm.Engine
}

func (od *OrderDao) CreateOrder(model Dto.OrderDto) (int64, error) {
	return od.engine.Insert(model)
}

func (od *OrderDao) QueryOrder(pageSearch Dto.GetOrderPageSearch) ([]Dto.GetOrderResult, error) {
	result := []Dto.GetOrderResult{}
	_, err := od.engine.Alias("o").Where("o.ordername like '%?%'", pageSearch.OrderName).Asc("orderid").Get(&result)
	return result, err
}
