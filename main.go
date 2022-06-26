package main

import (
	"cstungthanh/playground/data-structure/queue"
	"fmt"
)

func main() {
	q := queue.CreateQueueV1()
	fmt.Println("Init queue", q)

	q.Enqueue(1)
	fmt.Println("After enqueue", q, q.Lenght())

	q.Enqueue(2)
	fmt.Println("After enqueue", q, q.Lenght())

	q.Enqueue(3)
	fmt.Println("After enqueue", q, q.Lenght())

	v := q.Dequeue()
	fmt.Println("After dequeue", v, q)

	q.Clear()
	fmt.Println("After clear", q)
}
