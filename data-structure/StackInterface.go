package stack

type IStack interface {
	Push(element *ListNode)
	Pop() *ListNode
}
