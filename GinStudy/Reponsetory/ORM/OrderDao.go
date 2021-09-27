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
func (od *OrderDao) CreateOrder(model *Dto.Order) (int64, error) {
	return od.engine.Insert(model)
}

//更新订单
func (od *OrderDao) UpdateOrder(model *Dto.Order) (int64, error) {
	//传入的为结构体指针时，只有非空和0的field才会被作为更新的字段
	return od.engine.ID(model.OrderId).Update(model)
	//如果需要更新一个值为0，则此种方法将无法实现
	//affected, err := od.engine.Id(model.OrderId).Cols("age").Update(model)
}

//删除订单
func (od *OrderDao) DeleteOrder(model *Dto.Order) (int64, error) {
	return od.engine.ID(model.OrderId).Delete(model)
}

//查询订单(单表集合查询)
func (od *OrderDao) QueryOrder(pageSearch Dto.GetOrderPageSearch) ([]Dto.GetOrderResult, error) {
	result := make([]Dto.GetOrderResult, 0)
	//这个地方必须指定table，不然会用struct的名称作为表名
	session := od.engine.Cols("OrderId", "OrderName", "IsVirtual", "ShopId", "CreateTime")
	session.Table("Order").Alias("o").Where("1=1")
	if pageSearch.OrderName != "" {
		session.And("o.ordername like ?", "%"+pageSearch.OrderName+"%")
	}
	session.Desc("OrderId")
	err := session.Find(&result)
	return result, err
}

//=============================各种形式的sql查询=============================
//多表连接查询
func (od *OrderDao) QueryOrderDeatil(pageSearch Dto.GetOrderDetailPageSearch) ([]Dto.GetOrderDetailResult, error) {
	result := make([]Dto.GetOrderDetailResult, 0)
	session := od.engine.Cols("OrderDetailId", "d.OrderId", "d.OrderName", "IsVirtual", "ShopId", "d.CreateTime")
	session.Table("order").Alias("o").Join("LEFT", "OrderDetail d", "o.OrderId=d.OrderId").Where("1=1")
	if pageSearch.OrderName != "" {
		session.And("d.ordername like ?", "%"+pageSearch.OrderName+"%")
	}
	err := session.Find(&result)
	return result, err
}
