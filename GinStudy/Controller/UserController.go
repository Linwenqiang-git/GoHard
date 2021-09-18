package controller

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	Dto "github.linwenqiang.com/GinStudy/Model/Dto"

	//. "github.linwenqiang.com/GinStudy/Service/IService"
	service "github.linwenqiang.com/GinStudy/Service"
)

//控制器结构体
type UserController struct {
	_userService *service.UserService `inject:""`
}

/*======================================内部action 不对外提供使用======================================*/
//用户登录
func (this *UserController) Login(context *gin.Context) {
	//这一部分相当于前置处理
	var user Dto.UserDto
	err := context.Bind(&user)
	if err != nil {
		log.Fatal(err.Error())
		println("获取参数出错：" + err.Error())
	}
	fmt.Println("正常调用到conterller")
	result := this._userService.Login(user.Username, user.Password)

	//Writer的两种写法
	// if result {
	// 	context.Writer.Write([]byte("hello," + user.Username + " success login"))
	// } else {
	// 	context.Writer.WriteString("hello," + user.Password + " fail")
	// }
	context.JSON(200, &result)
}

//获取用户信息
func (this *UserController) GetUserInfo(context *gin.Context) {
	userId := context.Param("userid")
	id, _ := strconv.Atoi(userId)

	dataResult := this._userService.GetUserInfo(id)

	result := Dto.ResponseModel{
		Code:   200,
		ErrMsg: "",
		Data:   dataResult,
	}
	//json格式返回数据
	context.JSON(200, &result)
}
