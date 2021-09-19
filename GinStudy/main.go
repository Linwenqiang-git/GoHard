package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	. "github.linwenqiang.com/GinStudy/Controller"
	. "github.linwenqiang.com/GinStudy/DependentInjection"
	. "github.linwenqiang.com/GinStudy/MiddleWare"
	Dto "github.linwenqiang.com/GinStudy/Model/Dto"
	"go.uber.org/dig"
)

func main() {
	/*
		默认使用了Logger、Recovery两个中间件
		Recovery当程序出错时，恢复程序并默认返回500错误，类似于全局异常捕获
	*/
	engine := gin.Default()
	/*
		中间件调用
		1.全局请求使用；
		2.单个接口只用，可直接在第二个参数，把中间件方法传入
	*/
	engine.Use(RequestHandle())

	//依赖注入
	container := ServiceBaseConfiguration()

	//设置html目录
	//engine.LoadHTMLGlob("./html/*")
	//静态文件加载 第一个参数：前端请求路径， 第二个参数：项目文件路径
	//engine.Static("/img", "./img")

	//绑定路由
	BindingUserControllerRouting(engine, container)

	err := engine.Run(":8090")
	if err != nil {
		println("start server fail:", err)
	}
}

/*======================================内部action 不对外提供使用======================================*/
//相当于一个controller 绑定该conterller下面的action
func BindingUserControllerRouting(engine *gin.Engine, container *dig.Container) {
	var controlelrName = "/User"

	//路由分组，对应controller的划分
	UserRoute := engine.Group(controlelrName)
	{
		UserRoute.POST("/login", func(context *gin.Context) {
			//这一部分相当于前置处理
			var user Dto.UserDto
			err := context.Bind(&user)
			if err != nil {
				log.Fatal(err.Error())
				println("获取参数出错：" + err.Error())
			}
			Containererr := container.Invoke(Login)
			if err != nil {
				fmt.Println("调用异常")
				fmt.Println(Containererr)
			}

			context.JSON(200, "我成功返回啦")
		})
		UserRoute.GET("/GetUserInfo/:userid", func(context *gin.Context) {
			//PrintDI_Error(container.Provide(*context))
			PrintDI_Error(container.Invoke(GetUserInfo))
		})
	}
}
