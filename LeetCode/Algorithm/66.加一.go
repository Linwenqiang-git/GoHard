package algorithm

import "fmt"

func Call_66() {
	//digits := []int{9, 9, 9}
	//digits := []int{1, 9, 9}
	digits := []int{0}
	result := plusOne(digits)
	fmt.Print(result)
}

//官方解法，只需要关注末尾出现了几个9
func plusOne2(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	// digits 中所有的元素均为 9

	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}

func plusOne(digits []int) []int {
	length := len(digits)
	radix := 0
	for i := length - 1; i >= 0; i-- {
		value := digits[i]
		if i == length-1 {
			value += 1
		}
		if value+radix >= 10 {
			digits[i] = 0
			radix = 1
		} else {
			digits[i] = value + radix
			radix = 0
		}
	}
	if radix != 0 {
		result := make([]int, 0)
		result = append(result, 1)
		for i := 0; i < length; i++ {
			result = append(result, digits[i])
		}
		return result
	}
	return digits
}
