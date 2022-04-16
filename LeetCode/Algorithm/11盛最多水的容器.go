package algorithm

import (
	utility "github.linwenqiang.com/LeetCode/Utility"
)

//双指针，这种做法可行但是会超时
func maxArea(height []int) int {
	area := 0
	for left := 0; left < len(height)-1; left++ {
		for right := len(height) - 1; right > left; right-- {
			area = utility.Max(area, (right-left)*utility.Min(height[left], height[right]))
		}
	}
	return area
}

func maxArea2(height []int) int {
	area := 0
	left := 0
	right := len(height)
	for left < right {
		area = utility.Max(area, (right-left)*utility.Min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}
