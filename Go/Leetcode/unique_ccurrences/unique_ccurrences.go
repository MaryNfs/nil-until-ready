package unique_ccurrences

func uniqueOccurrences(arr []int) bool {
	hashmap := make(map[int]int, 0)
	res := make(map[int]int, 0)

	for i := 0; i < len(arr); i++ {
		hashmap[arr[i]]++
	}
	for _, v := range hashmap {
		_, is := res[v]
		if is {
			return false
		} else {
			res[v]++
		}
	}

	return true
}
