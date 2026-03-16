package guess_number

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int, pick int) int {
	if num > pick {
		return -1
	} else if num < pick {
		return 1
	}
	return 0
}

func guessNumber(n int, p int) int {
	l, r := 1, n
	for true {
		m := (l + r) / 2
		res := guess(m, p)
		if res == 0 {
			return m
		} else if res > 0 {
			l = m + 1
		} else if res < 0 {
			r = m - 1
		}
	}
	return 1
}
