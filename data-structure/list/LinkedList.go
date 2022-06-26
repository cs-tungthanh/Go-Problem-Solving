package list

type ListNode struct {
	Data int
	Next *ListNode
}

func CreateListNode(v int) *ListNode {
	return &ListNode{Data: v, Next: nil}
}
