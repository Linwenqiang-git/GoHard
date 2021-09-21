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

func (oc *OrderController) CreateOrder(context *gin.Context) {
	result := Dto.NewResult(context)
	//controller相当于前置处理
	var model Dto.OrderDto
	err := context.Bind(&model)
	if err != nil {
		log.Fatal(err.Error())
		println("获取参数出错：" + err.Error())
		result.Error(500, "获取参数出错")
	}
	model.CreateTime = time.Now()
	//调用具体的业务逻辑层
	service := impl.NewOrdeService()
	data := service.CreateOrder(model)
	if data > 0 {
		result.Success("创建订单成功")
	} else {
		result.Error(200, "创建订单失败:"+err.Error())
	}
}

/*======================================内部action 不对外提供使用======================================*/
//相当于一个controller 绑定该conterller下面的action
func (oc *OrderController) BindingOrderControllerRouting(engine *gin.Engine) {
	var controlelrName = "/Order"

	//路由分组，对应controller的划分
	UserRoute := engine.Group(controlelrName)
	{
		UserRoute.POST("/CreateOrder", oc.CreateOrder)

	}
}
