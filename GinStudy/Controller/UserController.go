package controller

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	. "github.linwenqiang.com/GinStudy/DependentInjection"
	Dto "github.linwenqiang.com/GinStudy/Model/Dto"

	service "github.linwenqiang.com/GinStudy/Service"
)

//控制器结构体
type UserController struct {
	//如果service没有添加到容器内的，需要dig.In显示注入，但是一定不能只指针类型
	dig.In
	UserService service.UserService
}

/*======================================内部action 不对外提供使用======================================*/
//用户登录
func Login(this UserController) {

	fmt.Println("正常调用到conterller")
	//result := this._userService.Login(user.Username, user.Password)
	this.UserService.Login("1", "2")
	//Writer的两种写法
	// if result {
	// 	context.Writer.Write([]byte("hello," + user.Username + " success login"))
	// } else {
	// 	context.Writer.WriteString("hello," + user.Password + " fail")
	// }

}

//获取用户信息
//无法正常调用，context 不能注入到dig容器中
func GetUserInfo(this UserController, context *gin.Context) {
	userId := context.Param("userid")
	fmt.Printf("成功获取到userid:%s", userId)
	id, _ := strconv.Atoi(userId)

	dataResult := this.UserService.GetUserInfo(id)
	result := Dto.NewResult(context)
	result.Success(dataResult)
}

/*======================================内部action 不对外提供使用======================================*/
//相当于一个controller 绑定该conterller下面的action
func BindingUserControllerRouting(engine *gin.Engine, container *dig.Container) {
	var controlelrName = "/User"

	//路由分组，对应controller的划分
	UserRoute := engine.Group(controlelrName)
	{
		UserRoute.POST("/login", func(context *gin.Context) {
			result := Dto.NewResult(context)
			//这一部分相当于前置处理
			var user Dto.UserDto
			err := context.Bind(&user)
			if err != nil {
				log.Fatal(err.Error())
				println("获取参数出错：" + err.Error())
				result.Error(500, "获取参数出错：")
			}
			PrintDI_Error(container.Invoke(Login))
			//无法获取到返回结果
			result.Success("我成功返回啦")
		})

		UserRoute.GET("/GetUserInfo/:userid", func(context *gin.Context) {
			controller := UserController{UserService: service.UserService{}}
			GetUserInfo(controller, context)
		})
	}
}
