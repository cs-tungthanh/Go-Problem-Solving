package queue

type QueueV1 struct {
	head int
	tail int
	data map[int]interface{}
}

func CreateQueueV1() *QueueV1 {
	return &QueueV1{head: 0, tail: 0, data: make(map[int]interface{}, 0)}
}

func (queue *QueueV1) Enqueue(data interface{}) {
	queue.data[queue.tail] = data
	queue.tail++
}

func (queue *QueueV1) Dequeue() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	if v, ok := queue.data[queue.head]; ok {
		delete(queue.data, queue.head)
		queue.head++
		return v
	}
	return nil
}

func (queue *QueueV1) Clear() {
	for k := range queue.data {
		delete(queue.data, k)
	}
	queue.head = 0
	queue.tail = 0
}

func (queue *QueueV1) IsEmpty() bool {
	return queue.Lenght() == 0
}

func (queue *QueueV1) Lenght() int {
	return queue.tail - queue.head
}
