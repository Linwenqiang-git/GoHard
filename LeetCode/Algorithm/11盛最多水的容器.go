package algorithm

import (
	utility "github.linwenqiang.com/LeetCode/Utility"
)

func maxArea(height []int) int {
	area := 0
	for left := 0; left < len(height)-1; left++ {
		for right := len(height) - 1; right > left; right-- {
			area = utility.Max(area, (right-left)*utility.Min(height[left], height[right]))
		}
	}
	return area
}
