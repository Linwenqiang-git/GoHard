package controller

import (
	"fmt"
	//"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	Dto "github.linwenqiang.com/GinStudy/Model/Dto"

	//. "github.linwenqiang.com/GinStudy/Service/IService"
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
func GetUserInfo(this UserController, context gin.Context) {
	userId := context.Param("userid")
	fmt.Printf("成功获取到userid:%s", userId)
	id, _ := strconv.Atoi(userId)

	dataResult := this.UserService.GetUserInfo(id)

	result := Dto.ResponseModel{
		Code:   200,
		ErrMsg: "",
		Data:   dataResult,
	}
	//json格式返回数据
	context.JSON(200, &result)
}
