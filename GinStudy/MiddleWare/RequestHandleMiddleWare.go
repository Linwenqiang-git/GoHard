package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RequestHandle() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method
		fmt.Println("============================Request MiddleWare============================")
		fmt.Println("请求路径：", path)
		fmt.Println("请求的方法：", method)

		/*
			相当于.net core 里的next RequestDelegate委托，俄罗斯套娃，将中间件一分为二
			区别在于，.net core里不调用next，后面的中间件就不会继续执行

		*/
		context.Next()

		fmt.Println("请求返回的状态码", context.Writer.Status())
	}
}
