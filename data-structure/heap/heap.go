package heap

func Heapify(nums []int, index int) {
	largest := index
	lChild := 2*index + 1
	rChild := 2*index + 2

	if lChild < len(nums) && nums[largest] < nums[lChild] {
		largest = lChild
	}
	if rChild < len(nums) && nums[largest] < nums[rChild] {
		largest = rChild
	}
	if largest != index {
		nums[largest], nums[index] = nums[index], nums[largest]
		Heapify(nums, largest)
	}
}

func HeapifyBottomUp(nums []int, index int) {
	parent := index / 2
	if parent >= 0 {
		if nums[parent] < nums[index] {
			nums[parent], nums[index] = nums[index], nums[parent]
			HeapifyBottomUp(nums, parent)
		}
	}
}

func PopHeap(nums *[]int) int {
	root := (*nums)[0]
	(*nums)[0] = (*nums)[len(*nums)-1]
	Heapify(*nums, 0)
	return root
}

func BuildHeap(nums []int) []int {
	// we need to iterate from the last index
	// because the heapify assuming that at index i it must be an correct heap
	for i := len(nums) - 1; i >= 0; i-- {
		Heapify(nums, i)
	}
	return nums
}

func InseartHeap(nums *[]int, value int) {
	*nums = append(*nums, value)
	HeapifyBottomUp(*nums, len(*nums)-1)
}
