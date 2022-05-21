package algorithm

import (
	"math"
	"sort"

	utility "github.linwenqiang.com/LeetCode/Utility"
)

//最接近平均值的两个值，计算其距离（方法不对）
func minMoves2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	sort.Ints(nums)
	total := 0
	for _, val := range nums {
		total += val
	}
	avg := total / len(nums)
	left, right := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= avg && nums[i+1] >= avg {
			left = nums[i]
			right = nums[i+1]
			break
		}
	}
	left_distance, right_distance := 0, 0
	for _, val := range nums {
		left_distance += int(math.Abs(float64(left) - float64(val)))
		right_distance += int(math.Abs(float64(right) - float64(val)))
	}
	return utility.Min(left_distance, right_distance)
}

//排序之后取中间位置的那个数，证明省略。。。
func minMoves22(nums []int) (ans int) {
    sort.Ints(nums)
    x := nums[len(nums)/2]
    for _, num := range nums {
        ans += abs(num - x)
    }
    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
