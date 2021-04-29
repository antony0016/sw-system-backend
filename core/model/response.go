package model

type Response struct {
	ErrorCode    int
	ErrorMessage string
	Message      string
	Data         interface{}
}
