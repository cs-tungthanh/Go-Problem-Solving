package stack

type IStack interface {
	Push(element interface{})
	Pop() interface{}
}
