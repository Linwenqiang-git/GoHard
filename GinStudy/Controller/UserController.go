package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	dtoModel "github.linwenqiang.com/GinStudy/Dto" //.表示调用公共方法不需要加前缀
)

var controlelrName = "/User"

//相当于action方法
func Login(user dtoModel.UserDto) bool {
	if user.Username == "1" && user.Password == "2" {
		return true
	} else {
		return false
	}
}

//相当于一个controller 绑定该conterller下面的action
func BindingUserControllerRouting(engine *gin.Engine) {

	engine.POST(controlelrName+"/login", func(context *gin.Context) {
		fmt.Println("参数 username:", context.PostForm("username"))
		//这一部分相当于前置处理
		var user dtoModel.UserDto
		err := context.BindQuery(user)
		if err != nil {
			log.Fatal(err.Error())
			println("获取参数出错：" + err.Error())
		}
		fmt.Println(" Username = ", user.Username)
		fmt.Println(" Password = ", user.Password)
		result := Login(user)
		if result {
			context.Writer.Write([]byte("hello," + user.Username + " success login"))
		} else {
			context.Writer.Write([]byte("hello," + user.Password + " fail"))
		}
	})
}
