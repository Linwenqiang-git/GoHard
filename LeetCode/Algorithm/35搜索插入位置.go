package algorithm

func searchInsert(nums []int, target int) int {
	//排序数组，最快是二分法查找
	middle := len(nums) / 2
	result := 0
	//从左侧找
	if target < nums[middle] {
		for i := 0; i < middle; i++ {
			if nums[i] >= target {
				return i
			}
		}
		result = middle //找不到就是最大的
	} else if target > nums[middle] {
		for i := middle + 1; i < len(nums); i++ {
			if nums[i] >= target {
				return i
			}
		}
		result = len(nums)
	} else {
		return middle
	}
	return result
}

func Call_35() {

}
