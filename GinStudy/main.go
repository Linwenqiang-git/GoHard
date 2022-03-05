package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	controller "github.linwenqiang.com/GinStudy/Controller"
	. "github.linwenqiang.com/GinStudy/DependentInjection"
	. "github.linwenqiang.com/GinStudy/MiddleWare"
	_ "github.linwenqiang.com/GinStudy/docs" // 执行swag init生成的docs文件夹路径 _引包表示只执行init函数
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
	engine.Use(Cors())

	// swagger访问地址 http://localhost:8080/swagger/index.html
	url := ginSwagger.URL("http://localhost:8090/swagger/doc.json") // The url pointing to API definition
	// 添加swagger的路由  不然报错404 page not found
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	//依赖注入
	container := ServiceBaseConfiguration()

	//设置html目录
	//engine.LoadHTMLGlob("./html/*")
	//静态文件加载 第一个参数：前端请求路径， 第二个参数：项目文件路径
	//engine.Static("/img", "./img")

	//绑定路由
	controller.BindingUserControllerRouting(engine, container)
	//这种路由绑定方式比较友好
	new(controller.OrderController).BindingOrderControllerRouting(engine)

	err := engine.Run("127.0.0.1:8090")
	if err != nil {
		println("start server fail:", err)
	}
}
