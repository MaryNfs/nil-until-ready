package max_score

import (
	"container/heap"
	"sort"
)

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	pairs := make([][2]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{nums2[i], nums1[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})

	h := &minHeap{}
	heap.Init(h)

	var sum1 int64 = 0
	var best int64 = 0

	for _, p := range pairs {
		cur2 := int64(p[0])
		cur1 := p[1]

		heap.Push(h, cur1)
		sum1 += int64(cur1)

		if h.Len() > k {
			sum1 -= int64(heap.Pop(h).(int))
		}

		if h.Len() == k {
			score := sum1 * cur2
			if score > best {
				best = score
			}
		}
	}

	return best
}

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}


// python solution
/*
class Solution:
    def maxScore(self, nums1: List[int], nums2: List[int], k: int) -> int:
        total = res = 0
        heap = []
        pairs = zip(nums1, nums2)
        sorted_pairs = sorted(pairs, key=lambda x: -x[1])

        for pair in sorted_pairs:
            num1, num2 = pair
            heappush(heap, num1)
            total += num1
            
            if len(heap) > k:
                total -= heappop(heap)
            if len(heap) == k:
                res = max(res, total * num2)

        return res
*/



