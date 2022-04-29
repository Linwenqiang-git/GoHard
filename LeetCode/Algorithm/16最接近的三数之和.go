package algorithm

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	result := math.MaxInt32
	for i := 0; i < n; i++ {
		Pb := i + 1
		Pc := n - 1
		for Pb < Pc {
			sum := nums[i] + nums[Pb] + nums[Pc]
			if sum == target {
				return target
			}

			cur := int(math.Abs(float64(sum - target)))
			if cur < int(math.Abs(float64(result-target))) {
				result = sum
			}
			if sum > target {
				//右指针左移
				Pc--
			} else {
				//左指针
				Pb++
			}
		}
	}
	return result
}

func Call_16() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4, 5}, 3))
}
