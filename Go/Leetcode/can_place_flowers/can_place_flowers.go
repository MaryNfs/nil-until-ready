package can_place_flowers

// The ugly
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	if n == 0 {
		return true
	}
	if len(flowerbed) == 1 {
		if flowerbed[0] == 0 {
			return true
		} else {
			return false
		}
	}
	for i := 0; i < len(flowerbed); i++ {
		if i == 0 {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				count++
				flowerbed[i] = 1
			}
		} else if i == len(flowerbed)-1 {
			if flowerbed[i] == 0 && flowerbed[i-1] == 0 {
				count++
				flowerbed[i] = 1
			}
		} else if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			count++
			flowerbed[i] = 1
		}
	}
	if count >= n {
		return true
	}
	return false
}
