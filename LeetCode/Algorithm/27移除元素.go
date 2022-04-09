package algorithm

import "fmt"

//后面取数组，就取返回int的位数，这样不会有额外的内存开销，执行速度快
func removeElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 0
	for fast := 0; fast < n; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func Call_27() {
	data := []int{3, 2, 2, 3}
	result := removeElement(data, 3)
	fmt.Printf("数组长度：%d\n", result)
	for _, val := range data {
		fmt.Printf("%d,", val)
	}
}
