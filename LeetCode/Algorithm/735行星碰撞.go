package algorithm

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, v := range asteroids {
		alive := true
		for alive && len(stack) > 0 && v < 0 && stack[len(stack)-1] > 0 {
			lenStack := len(stack)
			if -v > stack[lenStack-1] {
				stack = stack[:lenStack-1]
			} else if -v == stack[lenStack-1] {
				stack = stack[:lenStack-1]
				alive = false
			} else {
				alive = false
			}
		}
		if alive {
			stack = append(stack, v)
		}
	}
	return stack
}
