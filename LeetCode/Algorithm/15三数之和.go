package algorithm

import (
	"fmt"
	"sort"
	"strconv"
)

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) == 0 {
		return result
	}
	sort.Ints(nums)
	dic := make(map[string]bool)
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[i]+nums[j] <= 0 {
					for x := j + 1; x < len(nums); x++ {
						if nums[i]+nums[j]+nums[x] == 0 {
							sumstr := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(nums[x])
							if !dic[sumstr] {
								result = append(result, []int{nums[i], nums[j], nums[x]})
								dic[sumstr] = true
							}
							break
						}
						if nums[i]+nums[j]+nums[x] > 0 {
							break
						}
					}
				} else {
					break
				}
			}
		} else {
			break
		}
	}
	return result
}

//双指针的思路
func threeSum2(nums []int) [][]int {
	//先按照从小到大排序
	sort.Ints(nums)
	n := len(nums)
	result := make([][]int, 0)
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first - 1] {
            continue
        }
		//第二层和第三层使用双指针
		third := n - 1
		for second := first + 1; second < n; second++ {
			//需要和上一次枚举的数不相同
			if second == first+1 || nums[second] != nums[second-1] {
				if nums[first]+nums[second] <= 0 {
					for nums[first]+nums[second]+nums[third] > 0 && third > second {
						third--
					}
					if second == third {
						break
					}
					if nums[first]+nums[second]+nums[third] == 0 {
						result = append(result, []int{nums[first], nums[second], nums[third]})
					}
				}
			}
		}
	}
	return result
}

func Call_15() {
	fmt.Println(threeSum2([]int{1, 1, -2}))
}
