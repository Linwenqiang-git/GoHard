package service

import (
	"fmt"

	Dto "github.linwenqiang.com/GinStudy/Model/Dto"
	repo "github.linwenqiang.com/GinStudy/Reponsetory"
	"go.uber.org/dig"
)

type UserService struct {
	dig.In
	Repo repo.DbContext
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
func (user *UserService) GetUserInfo(id int) []interface{} {
	var context repo.IDbContext = &repo.DbContext{}
	context.SetConnect("root:LWQlwq123@@tcp(127.0.0.1:3306)/golangstudy")
	result := context.Query("select * from UserDto", &Dto.UserDto{})
	return result
}
