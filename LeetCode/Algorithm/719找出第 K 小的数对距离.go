package algorithm

import "sort"

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	return sort.Search(nums[len(nums)-1]-nums[0], func(mid int) bool {
		cnt, i := 0, 0
		for j, num := range nums {
			for num-nums[i] > mid {
				i++
			}
			cnt += j - i
		}
		return cnt >= k
	})
}
