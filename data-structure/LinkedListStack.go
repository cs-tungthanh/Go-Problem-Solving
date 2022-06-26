package stack

type StackNode struct {
	data *ListNode
	next *StackNode
}

func CreateStackNode(data *ListNode, next *StackNode) *StackNode {
	return &StackNode{data, next}
}

type StackV1 struct {
	top  *StackNode
	size int
}

func (stack *StackV1) Push(element *ListNode) {
	stack.top = CreateStackNode(element, stack.top)
	stack.size++
}

func (stack *StackV1) Pop() *ListNode {
	if stack.size > 0 {
		el := stack.top.data
		stack.top = stack.top.next
		stack.size--
		return el
	}
	return nil
}

func (stack *StackV1) GetSize() int {
	return stack.size
}

func (stack *StackV1) IsEmpty() bool {
	return stack.size == 0
}

func CreateLinkedListStack() IStack {
	stack := &StackV1{
		top:  nil,
		size: 0,
	}
	return stack
}
