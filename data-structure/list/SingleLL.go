package list

import "fmt"

type ListNode struct {
	Data interface{}
	Next *ListNode
}

func CreateListNode(v interface{}) *ListNode {
	return &ListNode{Data: v, Next: nil}
}

type SingleLL struct {
	head *ListNode
}

func CreateSingleLL() *SingleLL {
	return &SingleLL{nil}
}

func (list *SingleLL) AppendNode(data interface{}) {
	node := CreateListNode(data)
	if list.head == nil {
		list.head = node
	} else {
		curr := list.head
		list.head = node
		list.head.Next = curr
	}
}

func (list *SingleLL) Reverse() {
	curr := list.head
	var prev *ListNode
	var next *ListNode
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	list.head = prev
}

func (list *SingleLL) Display() {
	curr := list.head
	for curr != nil {
		fmt.Println(curr.Data)
		curr = curr.Next
	}
}

/*
	ll := list.CreateSingleLL()
	ll.AppendNode(1)
	ll.AppendNode(2)
	ll.AppendNode(3)
	ll.Display()
	ll.Reverse()
	ll.Display()
*/
