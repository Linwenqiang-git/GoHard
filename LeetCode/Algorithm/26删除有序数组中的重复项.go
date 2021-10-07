package main

import "fmt"

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums); {
		//fmt.Printf("i=%d\n", i)
		if i+1 < len(nums) {
			if nums[i] == nums[i+1] {
				nums = append(nums[:i], nums[i+1:]...)
			} else {
				i++
			}
		} else {
			break
		}
	}
	return len(nums)
}

//官方解法
func removeDuplicates1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main() {
	data := []int{1, 1, 1, 2, 3, 4}
	result := removeDuplicates1(data)
	fmt.Printf("数组长度：%d\n", result)
	for _, val := range data {
		fmt.Printf("%d,", val)
	}
}
