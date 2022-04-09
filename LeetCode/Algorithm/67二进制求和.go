package algorithm

import (
	"fmt"
	"strconv"
)

func Call_67() {
	data := 1 / 2
	fmt.Println(data)
	fmt.Print(addBinary("1011", "1010"))
}

func addBinary(a string, b string) string {
	var radix int = 0
	i := len(a) - 1
	j := len(b) - 1

	result := ""
	for i >= 0 || j >= 0 {
		var ai, bj int = 0, 0
		if i >= 0 {
			ai = int(a[i] - '0')
		}
		if j >= 0 {
			bj = int(b[j] - '0')
		}
		value := ai + bj + radix
		if value >= 2 {
			result = strconv.Itoa(value%2) + result
			radix = 1
		} else {
			result = strconv.Itoa(value) + result
			radix = 0
		}
		i -= 1
		j -= 1
	}
	if radix > 0 {
		result = "1" + result
	}
	return result
}

//官方解法，巧妙的运行进制位
func addBinary2(a string, b string) string {
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}
