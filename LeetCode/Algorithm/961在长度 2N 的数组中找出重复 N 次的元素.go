package algorithm

import "sort"

func repeatedNTimes(nums []int) int {
	sort.Ints(nums)
	ret := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[i+1] {
			ret = nums[i]
			break
		}
	}
	return ret
}
