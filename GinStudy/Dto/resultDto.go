package dto

type ResponseModel struct {
	Code   int
	ErrMsg string
	Data   interface{}
}
