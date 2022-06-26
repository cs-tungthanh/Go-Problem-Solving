package main

import (
	"cstungthanh/playground/data-structure/list"
)

func main() {
	ll := list.CreateSingleLL()
	ll.AppendNode(1)
	ll.AppendNode(2)
	ll.AppendNode(3)
	ll.Display()

	ll.Reverse()
	ll.Display()
}
