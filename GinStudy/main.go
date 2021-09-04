package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	//. "github.linwenqiang.com/GinStudy/Controller"
)

type UserDto struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func Login(user UserDto) bool {
	if user.Username == "1" && user.Password == "2" {
		return true
	} else {
		return false
	}
}
func main() {
	engine := gin.Default()

	//BindingUserControllerRouting(engine)

	engine.POST("/User/login", func(context *gin.Context) {
		fmt.Println("参数 username:", context.PostForm("username"))
		//这一部分相当于前置处理
		var user UserDto
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

	err := engine.Run(":8090")
	if err != nil {
		println("start server fail:", err)
	}
}
