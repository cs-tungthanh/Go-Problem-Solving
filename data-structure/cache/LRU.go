package cache

type DoubleNode struct {
	val  int
	key  int
	Next *DoubleNode
	Prev *DoubleNode
}

func CreateNode(k, v int) *DoubleNode {
	return &DoubleNode{
		val:  v,
		key:  k,
		Next: nil,
		Prev: nil,
	}
}

type LRUCache struct {
	csize int
	len   int
	head  *DoubleNode
	tail  *DoubleNode
	dict  map[int]*DoubleNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		csize: capacity,
		len:   0,
		head:  nil,
		tail:  nil,
		dict:  make(map[int]*DoubleNode),
	}
}

func (cache *LRUCache) removeNode(node *DoubleNode) {
	// remove link
	if node == cache.head {
		cache.head = node.Next
	} else if node == cache.tail {
		cache.tail = node.Prev
	} else {
		currNext := node.Next
		currPrev := node.Prev
		// connect prev and next
		if currNext != nil {
			currNext.Prev = currPrev
		}
		if currPrev != nil {
			currPrev.Next = currNext
		}
	}
}

func (cache *LRUCache) Get(key int) int {
	node, ok := cache.dict[key]
	if ok {
		if node != cache.tail {
			cache.removeNode(node)
			// append to the end
			cache.tail.Next = node
			node.Prev = cache.tail
			node.Next = nil
			cache.tail = node
		}
		return node.val
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	node := CreateNode(key, value)
	oldNode, ok := cache.dict[key]
	if ok {
		// remove old node
		cache.len -= 1
		cache.removeNode(oldNode)
	} else if cache.len == cache.csize {
		// evict
		cache.len -= 1
		delete(cache.dict, cache.head.key)
		if cache.head == cache.tail {
			cache.tail = cache.tail.Next
		}
		cache.head = cache.head.Next
	}

	cache.len += 1
	cache.dict[key] = node
	if cache.len == 1 {
		cache.head = node
		cache.tail = node
		return
	}
	cache.tail.Next = node
	node.Prev = cache.tail
	cache.tail = node
}
