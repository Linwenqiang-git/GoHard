package algorithm

func findDuplicates(nums []int) []int {
	ret := []int{}
	if len(nums) == 0{
		return ret
	}
	for _, x := range nums {
        if x < 0 {
            x = -x
        }
        if nums[x-1] > 0 {
            nums[x-1] = - nums[x-1]
        } else {
            ret = append(ret, x)
        }
    }
    return ret
}
