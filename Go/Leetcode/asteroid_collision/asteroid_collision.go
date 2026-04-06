package asteroid_collision

func asteroidCollision(asteroids []int) []int {
	var stack []int
	for i := 0; i < len(asteroids); i++ {
		for len(stack) > 0 && asteroids[i] < 0 && stack[len(stack)-1] > 0 {
			diff := asteroids[i] + stack[len(stack)-1]
			if diff < 0 {
				stack = stack[:len(stack)-1]
			} else if diff > 0 {
				asteroids[i] = 0
			} else {
				asteroids[i] = 0
				stack = stack[:len(stack)-1]
			}
		}
		if asteroids[i] != 0 {
			stack = append(stack, asteroids[i])
		}
	}
	return stack
}