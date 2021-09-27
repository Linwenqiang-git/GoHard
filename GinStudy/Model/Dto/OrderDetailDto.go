package dto

import "time"

//订单明细表
type OrderDetail struct {
	OrderDetailId int       `xorm:"pk autoincr 'OrderId'"`
	OrderId       int       `xorm:"bigint 'OrderId'"`
	OrderName     string    `xorm:"varchar(100) 'OrderName'" form:"ordername"`
	IsVirtual     byte      `xorm:"tinyint 'IsVirtual'" form:"isVirtual"`
	ShopId        int       `xorm:"bigint 'ShopId'" form:"shopId"`
	CreateTime    time.Time `xorm:"datetime 'CreateTime'" form:"createTime"`
}

//===============================订单查询Dto===============================
type GetOrderDetailPageSearch struct {
	OrderName string
}

type GetOrderDetailResult struct {
	OrderDetailId int       `xorm:"'OrderDetailId'"`
	OrderId       int       `xorm:"'OrderId'"`
	OrderName     string    `xorm:"'OrderName'"`
	IsVirtual     byte      `xorm:"'IsVirtual'"`
	ShopId        int       `xorm:"'ShopId'"`
	CreateTime    time.Time `xorm:"'CreateTime'"`
}

//===============================订单查询Dto===============================
