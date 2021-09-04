package main

import (
	"github.com/gin-gonic/gin"
	. "github.linwenqiang.com/GinStudy/Controller"
)

func main() {
	engine := gin.Default()

	BindingUserControllerRouting(engine)

	err := engine.Run(":8090")
	if err != nil {
		println("start server fail:", err)
	}
}
