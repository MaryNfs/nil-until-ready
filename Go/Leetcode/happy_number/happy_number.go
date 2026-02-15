package happy_number

var check = make(map[int]int, 0)

func isHappy(n int) bool {
	if n == 1 {
		check = make(map[int]int)
		return true
	}
	numArr := NumToArrayMath(n)
	res := 0
	for i := 0; i < len(numArr); i++ {
		res += numArr[i] * numArr[i]
	}
	_, ok := check[res]
	if ok != true {
		check[res] = 1
		return isHappy(res)
	} else {
		check = make(map[int]int)
		return false
	}

}

func NumToArrayMath(num int) []int {
	var arr []int
	if num == 0 {
		return []int{0}
	}
	for num > 0 {
		digit := num % 10
		arr = append(arr, digit) // Digits are appended in reverse order
		num /= 10
	}

	// Reverse the slice to correct the order
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
