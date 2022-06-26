package main

import (
	ds "cstungthanh/playground/data-structure"
	"fmt"
)

func main() {
	st := ds.CreateLinkedListStack()
	st.Push(ds.CreateListNode(3))
	fmt.Println(st)

	v := st.Pop()
	fmt.Println(v)
}
