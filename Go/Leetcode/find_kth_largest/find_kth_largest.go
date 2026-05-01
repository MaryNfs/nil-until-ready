package find_kth_largest

import (
	"container/heap"
)

// ref: https://www.youtube.com/watch?v=ZmGk7h8KZLs
// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {

	// since it's a min heap, we should add minus
	for k, _ := range nums {
		nums[k] = -nums[k]
	}

	h := (*IntHeap)(&nums)

	// 1. Initialize the heap (reorders elements into a heap)
	heap.Init(h)

	for range k - 1 {
		heap.Pop(h)
	}

	return -(*h)[0]
}
