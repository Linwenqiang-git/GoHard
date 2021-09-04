package entity

type UserDto struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
