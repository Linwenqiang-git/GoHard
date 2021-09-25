package dto

import "time"

type Order struct {
	OrderId    int       `xorm:"pk autoincr 'OrderId'"`
	OrderName  string    `xorm:"varchar(100) 'OrderName'" form:"ordername"`
	IsVirtual  byte      `xorm:"tinyint 'IsVirtual'" form:"isVirtual"`
	ShopId     int       `xorm:"bigint 'ShopId'" form:"shopId"`
	CreateTime time.Time `xorm:"datetime 'CreateTime'" form:"createTime"`
}

//===============================订单查询Dto===============================
type GetOrderPageSearch struct {
	OrderName string
}

type GetOrderResult struct {
	OrderId    int       `xorm:"pk autoincr 'OrderId'"`
	OrderName  string    `xorm:"varchar(100) 'OrderName'" form:"ordername"`
	IsVirtual  byte      `xorm:"tinyint 'IsVirtual'" form:"isVirtual"`
	ShopId     int       `xorm:"bigint 'ShopId'" form:"shopId"`
	CreateTime time.Time `xorm:"datetime 'CreateTime'" form:"createTime"`
}

//===============================订单查询Dto===============================
