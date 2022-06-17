package algorithm

import (
	"sort"
)

func findPairs(nums []int, k int) int {
	sort.Ints(nums)
	max := nums[len(nums)-1] - nums[0]
	n := len(nums)
	if k > max {
		return 0
	}
	i, j := 0, 1
	record := make(map[int]bool)
	for i < n && j < n {
		if i == j {
			j++
			continue
		}
		if abs_532(nums[j], nums[i]) < k {
			j++
		} else if abs_532(nums[j], nums[i]) > k {
			i++
		} else if abs_532(nums[j], nums[i]) == k {
			_, ok := record[nums[j]+nums[i]]
			if !ok {
				record[nums[j]+nums[i]] = true
			}
			if i < j {
				i++
			} else {
				j++
			}
		}
	}
	return len(record)
}
func abs_532(i, j int) int {
	if i >= j {
		return i - j
	} else {
		return j - i
	}
}
