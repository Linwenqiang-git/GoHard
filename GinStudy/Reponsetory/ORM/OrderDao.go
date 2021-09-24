package orm

import (
	"github.com/go-xorm/xorm"
	Dto "github.linwenqiang.com/GinStudy/Model/Dto"
)

//为orderService提供数据库访问层
type OrderDao struct {
	engine *xorm.Engine
}

//新增订单
func (od *OrderDao) CreateOrder(model Dto.Order) (int64, error) {
	return od.engine.Insert(model)
}

func (od *OrderDao) QueryOrder(pageSearch Dto.GetOrderPageSearch) ([]Dto.GetOrderResult, error) {
	result := []Dto.GetOrderResult{}
	//_, err := od.engine.Table("Order").Alias("o").Where("o.ordername = '?'", pageSearch.OrderName).Asc("orderid").Get(&result)
	_, err := od.engine.Table("Order").Alias("o").Asc("orderid").Rows(result)
	return result, err
}
