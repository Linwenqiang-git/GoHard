package dependentinjection

/*====================================
使用感受：
资料少，官方提示的demo过于简单
报错信息资料少，无法解决
*/
import (
	"database/sql"
	"log"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"

	//controller "github.linwenqiang.com/GinStudy/Controller"

	//repository "github.linwenqiang.com/GinStudy/InitSql"
	service "github.linwenqiang.com/GinStudy/Service"
)

func ConfigureUserServiceDI(app *gin.Engine) {
	//inject declare
	_, err := sql.Open("mysql", "root:LWQlwq123@@tcp(127.0.0.1:3306)/golangstudy")
	if err != nil {
		panic(err.Error())
	}
	//userController := controller.UserController{}
	//repo := repository.DbContext{Context: db}
	var userservice service.UserService
	//_ := repo.DbContext
	//Injection
	var injector inject.Graph
	err = injector.Provide(
		//&inject.Object{Value: &userController},
		&inject.Object{Name: "_userService", Value: &userservice},
		//inject.Object{Name: "Repo", Value: &repo},
	)
	if err != nil {
		log.Fatal("inject fatal 123: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("inject fatal 345: ", err)
	}

	//BindingUserControllerRouting(app, userController)
}

/*======================================内部action 不对外提供使用======================================*/
//相当于一个controller 绑定该conterller下面的action
// func BindingUserControllerRouting(engine *gin.Engine, userController controller.UserController) {
// 	var controlelrName = "/User"

// 	//路由分组，对应controller的划分
// 	UserRoute := engine.Group(controlelrName)
// 	{
// 		UserRoute.POST("/login", userController.Login)
// 		UserRoute.GET("/GetUserInfo/:userid", userController.GetUserInfo)
// 	}
// }
