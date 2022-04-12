package dynamicprogramming

//f(x) = f(x-1) + f(x-2)
//递归到x=44会超时
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

//官方方法，滚动数组思想
func climbStairs2(n int) int {
	//对于f(1) = 1, 保证后两位相加为1，所以用 0，1 前面的0用于补位
	x, y, z := 0, 0, 1
	for i := 1; i <= n; i++ {
		x = y
		y = z
		z = x + y
	}
	return z
}

func main() {
	climbStairs(1)
}
