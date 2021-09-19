package service

import (
	//"fmt"
	//"reflect"

	"fmt"

	. "github.linwenqiang.com/GinStudy/InitSql"
	Dto "github.linwenqiang.com/GinStudy/Model/Dto"
	"go.uber.org/dig"
)

type UserService struct {
	dig.In
	Repo DbContext
}

func (user *UserService) Login(username, password string) []interface{} {
	//var context IDbContext = &DbContext{ConnStr: "root:LWQlwq123@@tcp(127.0.0.1:3306)/golangstudy"}
	// if reflect.DeepEqual(user.Repo, DbContext{}) {
	// 	fmt.Println("Structure is empty")
	// }
	fmt.Println("正常调用到service")
	fmt.Println(user.Repo.Query("select * from UserDto", &Dto.UserDto{}))
	return nil
	//return user.Repo.Query("select * from UserDto", &Dto.UserDto{})
}
func (user *UserService) GetUserInfo(id int) interface{} {
	return &Dto.UserDto{
		Username: "1",
		Password: "2",
	}
}
