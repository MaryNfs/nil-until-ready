package top_K_frequent

import "sort"

type KeyValuePair struct {
	Key   int
	Value int
}

func topKFrequent(nums []int, k int) []int {
	hres := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		hres[nums[i]]++
	}
	var sortedByValue []KeyValuePair
	for k, v := range hres {
		sortedByValue = append(sortedByValue, KeyValuePair{k, v})
	}
	sort.Slice(sortedByValue, func(i, j int) bool {
		return sortedByValue[i].Value > sortedByValue[j].Value
	})
	c := 1
	var res []int
	for _, v := range sortedByValue {
		if c > k {
			break
		}
		res = append(res, v.Key)
		c++
	}
	return res
}
