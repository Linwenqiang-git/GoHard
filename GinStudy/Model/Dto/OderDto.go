package dto

import "time"

type OrderDto struct {
	OrderId    int       `xorm:"pk autoincr"`
	OrderName  string    `xorm:"varchar(100)" form:"ordername"`
	IsVirtual  byte      `xorm:"tinyint" form:"isVirtual"`
	ShopId     int       `xorm:"bigint" form:"shopId"`
	CreateTime time.Time `xorm:"datetime" form:"createTime"`
}
