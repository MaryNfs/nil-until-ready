package happy_number

// this is not the best solution, becuase the variable is defined on package level
var check = make(map[int]int, 0)

func isHappy_2(n int) bool {
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


