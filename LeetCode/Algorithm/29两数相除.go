package algorithm

import (
	"fmt"
	"math"
)

/*
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
返回被除数 dividend 除以除数 divisor 得到的商。

解法：
需要考虑被除数临界值的情况
*/

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}
	// 一般情况，使用二分查找
	// 将所有的正数取相反数，这样就只需要考虑一种情况
	rev := false
	if dividend > 0 {
		dividend = -dividend
		rev = !rev
	}
	if divisor > 0 {
		divisor = -divisor
		rev = !rev
	}
	ans := 0
	left, right := 1, math.MaxInt32
	for left <= right {
		rightShift := (right - left) >> 1
		mid := left + rightShift // 注意溢出，并且不能使用除法
		if quickAdd(divisor, mid, dividend) {
			ans = mid
			if mid == math.MaxInt32 { // 注意溢出
				break
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if rev {
		return -ans
	}
	return ans

}

// 快速乘
// x 和 y 是负数，z 是正数
// 判断 z * y >= x 是否成立
func quickAdd(y, z, x int) bool {
	for result, add := 0, y; z > 0; z >>= 1 { // 不能使用除法
		if z&1 > 0 {
			// 需要保证 result + add >= x
			if result < x-add {
				return false
			}
			result += add
		}
		if z != 1 {
			// 需要保证 add + add >= x
			if add < x-add {
				return false
			}
			add += add
		}
	}
	return true
}

func Call_29() {
	fmt.Println(divide(12, 5))
}
