package dto

type ResponseModel struct {
	Code   int
	ErrMsg string
	Count  int
	Data   interface{}
}
