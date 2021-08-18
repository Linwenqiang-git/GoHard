package datalogic

import "fmt"

//首字母大写 表示外面可访问这个方法
//罗马数字转换
func RomanToInt(s string) (result int) {
	dic := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000, "IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
	result = 0
	var next string
	for i := 0; i < len(s); i++ {
		print(" i = ", i)
		if i+1 < len(s) {
			next = string(s[i]) + string(s[i+1])

			value, ok := dic[next]
			if ok {
				result += value
				i++
				print("next = ", next)
				continue
			}
		}
		result += dic[string(s[i])]
	}
	return result
}

//两数之和
func TwoSum(nums []int, target int) (result []int) {
	var dict map[int]int = make(map[int]int)
	for x := 0; x < len(nums); x++ {
		middle := target - nums[x]
		value, ok := dict[middle]
		if ok {
			result = append(result, x)
			result = append(result, value)
			break
		} else {
			dict[nums[x]] = x
		}

	}
	return result
}

//回文数
func IsPalindrome(x int) bool {
	initx := x
	if x < 0 {
		return false
	}
	result := 0
	for x != 0 {
		digint := x % 10
		result = result*10 + digint
		x /= 10
	}
	if result == initx {
		return true
	} else {
		return false
	}
}

//冒泡排序
func Sort(nums []int) []int {

	numsLenth := len(nums)
	for i := 0; i < numsLenth; i++ {
		for j := 0; j < numsLenth-i; j++ {
			if j+1 < numsLenth-i {
				if nums[j] > nums[j+1] {
					middle := nums[j+1]
					nums[j+1] = nums[j]
					nums[j] = middle
				}
			}
		}
		//fmt.Printf("%v,", nums)
	}
	return nums
}

//初始化方法 -- 构造函数类似  如果有参数，比如依赖注入怎么处理？
func init() {
	fmt.Println("init datalogic...")
}
