package dto

type Dto interface {
	GetDto(interface{}) interface{}
	GetModel(interface{}) interface{}
}
