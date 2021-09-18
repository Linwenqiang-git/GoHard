package iservice

type IUserService interface {
	Login(username, password string) []interface{}
	GetUserInfo(id int) interface{}
}
