package utility //utility "github.linwenqiang.com/LeetCode/Utility"

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
