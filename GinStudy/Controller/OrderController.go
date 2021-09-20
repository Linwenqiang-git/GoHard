package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	Dto "github.linwenqiang.com/GinStudy/Model/Dto"
)

type OrderController struct {
}

func (oc *OrderController) CreateOrder(context *gin.Context) {
	result := Dto.NewResult(context)
	//这一部分相当于前置处理
	var model Dto.OrderDto
	err := context.Bind(&model)
	if err != nil {
		log.Fatal(err.Error())
		println("获取参数出错：" + err.Error())
		result.Error(500, "获取参数出错")
	}
	//调用具体的业务逻辑层
	service := impl.NewMemberService()
	result.Success("我成功返回啦")
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
