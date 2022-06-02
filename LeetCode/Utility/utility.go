package utility //utility "github.linwenqiang.com/LeetCode/Utility"
import "math"

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

//二进制转十进制
func BinaryToDecimal(binarry string) float64 {
	sum := 0.0
	n := len(binarry) - 1
	for i := n; i >= 0; i-- {
		if binarry[i] == '1' {
			sum += math.Pow(2, float64(n-i))
		}
	}
	return sum
}
