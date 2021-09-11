package main

import (
	"github.com/gin-gonic/gin"
	. "github.linwenqiang.com/GinStudy/Controller" //.表示调用公共方法不需要加前缀
	. "github.linwenqiang.com/GinStudy/MiddleWare"
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

	//设置html目录
	//engine.LoadHTMLGlob("./html/*")
	//静态文件加载 第一个参数：前端请求路径， 第二个参数：项目文件路径
	//engine.Static("/img", "./img")

	//绑定各个模块
	BindingUserControllerRouting(engine)

	err := engine.Run(":8090")
	if err != nil {
		println("start server fail:", err)
	}
}