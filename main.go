package main

import (
	"cstungthanh/playground/data-structure/list"
	"cstungthanh/playground/data-structure/stack"
	"fmt"
)

func main() {
	st := stack.CreateLinkedListStack()
	st.Push(list.CreateListNode(3))
	st.Push(4)
	fmt.Println(st)
	v := st.Pop()
	fmt.Println(v)
	v = st.Pop()
	fmt.Println(v)
}
