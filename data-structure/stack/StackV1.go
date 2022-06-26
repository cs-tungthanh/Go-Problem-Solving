package stack

type StackNode struct {
	data interface{}
	next *StackNode
}

func CreateStackNode(data interface{}, next *StackNode) *StackNode {
	return &StackNode{data, next}
}

type StackV1 struct {
	top  *StackNode
	size int
}

func (stack *StackV1) Push(element interface{}) {
	stack.top = CreateStackNode(element, stack.top)
	stack.size++
}

func (stack *StackV1) Pop() interface{} {
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

/*
	st := ds.CreateLinkedListStack()
	st.Push(ds.CreateListNode(3))
	st.Push(4)
	fmt.Println(st)
	v := st.Pop()
	fmt.Println(v)
	v = st.Pop()
	fmt.Println(v)
*/
