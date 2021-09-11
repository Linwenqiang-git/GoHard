package controller

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	Dto "github.linwenqiang.com/GinStudy/Dto"
	Entity "github.linwenqiang.com/GinStudy/Entity"
	. "github.linwenqiang.com/GinStudy/InitSql"
)

var controlelrName = "/User"

/*======================================内部action 不对外提供使用======================================*/
//用户登录
func login(user Dto.UserDto) bool {
	var context IDbContext = &DbContext{ConnStr: "root:LWQlwq123@@tcp(127.0.0.1:3306)/golangstudy"}
	context.SetDbConnect()
	result := context.Query("select * from UserDto", &Dto.UserDto{})
	for user := range result {
		fmt.Printf("%v", user)
	}

	if user.Username == "1" && user.Password == "2" {
		return true
	} else {
		return false
	}
}

//获取用户信息
func getUserInfo(userId int) Entity.UserEntity {
	entity := Entity.UserEntity{
		UserId:    1,
		UserName:  "张三",
		UserPhoto: "",
	}
	return entity
}

/*======================================内部action 不对外提供使用======================================*/

//相当于一个controller 绑定该conterller下面的action
func BindingUserControllerRouting(engine *gin.Engine) {

	//路由分组，对应controller的划分
	UserRoute := engine.Group(controlelrName)

	UserRoute.POST("/login", func(context *gin.Context) {
		//这一部分相当于前置处理
		var user Dto.UserDto
		err := context.Bind(&user)
		if err != nil {
			log.Fatal(err.Error())
			println("获取参数出错：" + err.Error())
		}

		result := login(user)

		//Writer的两种写法
		if result {
			context.Writer.Write([]byte("hello," + user.Username + " success login"))
		} else {
			context.Writer.WriteString("hello," + user.Password + " fail")
		}
	})

	UserRoute.GET("/GetUserInfo/:userid", func(context *gin.Context) {
		//这一部分相当于前置处理
		userId := context.Param("userid")
		id, _ := strconv.Atoi(userId)

		dataResult := getUserInfo(id)

		result := Dto.ResponseModel{
			Code:   200,
			ErrMsg: "",
			Data:   dataResult,
		}
		//json格式返回数据
		context.JSON(200, &result)
	})

}
