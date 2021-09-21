package dto

import "time"

type OrderDto struct {
	OrderId    int       `xorm:"pk autoincr"`
	OrderName  string    `xorm:"varchar(100)" form:"ordername"`
	IsVirtual  byte      `xorm:"tinyint" form:"isVirtual"`
	ShopId     int       `xorm:"bigint" form:"shopId"`
	CreateTime time.Time `xorm:"datetime" form:"createTime"`
}

//===============================订单查询Dto===============================
type GetOrderPageSearch struct {
	OrderName string
}

type GetOrderResult struct {
	OrderId    int
	OrderName  string
	IsVirtual  byte
	ShopId     int
	CreateTime time.Time
}

//===============================订单查询Dto===============================
