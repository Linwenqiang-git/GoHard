package algorithm


func wiggleSort(nums []int) {
	bucket := [5001]int{}
	//将每个value放入桶中，统计每个value出现的次数
	for _, value := range nums {
		bucket[value]++
	}
	j := 5000
	//填充大元素(从大到小排序 填充完大元素)
	for i := 1; i < len(nums); i += 2 {
		//中间可能会有空的元素
		for bucket[j] == 0 {
			j--
		}
		nums[i] = j
		bucket[j]--
	}
	//填充小元素
	for i := 0; i < len(nums); i += 2 {
		for bucket[j] == 0 {
			j--
		}
		nums[i] = j
		bucket[j]--
	}
}
