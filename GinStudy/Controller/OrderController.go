package controller

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	Dto "github.linwenqiang.com/GinStudy/Model/Dto"
	impl "github.linwenqiang.com/GinStudy/Service"
)

type OrderController struct {
}

//创建订单
func (oc *OrderController) CreateOrder(context *gin.Context) {
	result := Dto.NewResult(context)
	//controller相当于前置处理
	var model Dto.Order
	err := context.Bind(&model)
	if err != nil {
		log.Fatal(err.Error())
		println("获取参数出错：" + err.Error())
		result.Error(500, "获取参数出错")
	}
	model.CreateTime = time.Now()
	//调用具体的业务逻辑层
	service := impl.NewOrdeService()
	data := service.CreateOrder(&model)
	result.Success(data)
}

//更新订单
func (oc *OrderController) UpdateOrder(context *gin.Context) {
	result := Dto.NewResult(context)
	//controller相当于前置处理
	var model Dto.Order
	err := context.Bind(&model)
	if err != nil {
		log.Fatal(err.Error())
		println("获取参数出错：" + err.Error())
		result.Error(500, "获取参数出错")
	}
	model.CreateTime = time.Now()
	//调用具体的业务逻辑层
	service := impl.NewOrdeService()
	data := service.UpdateOrder(&model)
	result.Success(data)
}

//查询订单
func (oc *OrderController) GetOrder(context *gin.Context) {
	result := Dto.NewResult(context)
	var pageSearch Dto.GetOrderPageSearch
	err := context.Bind(&pageSearch)
	if err != nil {
		log.Fatal(err.Error())
		println("获取参数出错：" + err.Error())
		result.Error(500, "获取参数出错")
	}
	//调用具体的业务逻辑层
	service := impl.NewOrdeService()
	data := service.QueryOrder(pageSearch)
	result.Success(data)
}

//查询订单明细
func (oc *OrderController) GetOrderDetail(context *gin.Context) {
	result := Dto.NewResult(context)
	var pageSearch Dto.GetOrderDetailPageSearch
	err := context.Bind(&pageSearch)
	if err != nil {
		log.Fatal(err.Error())
		println("获取参数出错：" + err.Error())
		result.Error(500, "获取参数出错")
	}
	//调用具体的业务逻辑层
	service := impl.NewOrdeService()
	data := service.QueryOrderDeatil(pageSearch)
	result.Success(data)
}

/*======================================内部action 不对外提供使用======================================*/
//相当于一个controller 绑定该conterller下面的action
func (oc *OrderController) BindingOrderControllerRouting(engine *gin.Engine) {
	var controlelrName = "/Order"

	//路由分组，对应controller的划分
	UserRoute := engine.Group(controlelrName)
	{
		UserRoute.POST("/UpdateOrder", oc.UpdateOrder)
		UserRoute.POST("/CreateOrder", oc.CreateOrder)
		UserRoute.POST("/GetOrder", oc.GetOrder)
		UserRoute.POST("/GetOrderDetail", oc.GetOrderDetail)
	}
}
