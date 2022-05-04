package algorithm

import "fmt"

func search(nums []int, target int) int {
	n := len(nums)

	ret := -1
	left, right := 0, n-1
	for left <= right {
		middle := (right + left) / 2
		if nums[middle] == target {
			return middle
		}
		if nums[0] <= nums[middle] { //保证这一侧是顺序的
			if nums[0] <= target && target < nums[middle] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else {
			//那么这一侧是顺序的
			if nums[middle] < target && target <= nums[n-1] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}

	return ret
}

func Call_33() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}
