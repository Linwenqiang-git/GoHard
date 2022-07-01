package tree

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := []int{}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {

		p := []*TreeNode{}
		currentMax := -1 << 31
		for _, node := range q {
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
			if node.Val > currentMax {
				currentMax = node.Val
			}
		}
		ret = append(ret, currentMax)
		q = p
	}
	return ret
}

func majorityElement(nums []int) int {
	numCount := make(map[int]int)
	n := len(nums) / 2
	for _, value := range nums {
		numCount[value]++
		if numCount[value] > n {
			return value
		}
	}
	return nums[0]
}
