package algorithm

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	result := make([][]int, 0)
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first + 1; second < n; second++ {
			four := n - 1
			if second-first > 1 && nums[second] == nums[second-1] {
				continue
			}
			for third := second + 1; third < n; third++ {
				if third-second > 1 && nums[third] == nums[third-1] {
					continue
				}
				for nums[first]+nums[second]+nums[third]+nums[four] > target && four > third {
					four--
				}
				if third == four {
					break
				}
				if nums[first]+nums[second]+nums[third]+nums[four] == target {
					result = append(result, []int{nums[first], nums[second], nums[third], nums[four]})
				}
			}
		}
	}
	return result
}

func Call_18() {
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 0))
}
