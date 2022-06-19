package tree

func findFrequentTreeSum(root *TreeNode) []int {
	total := make(map[int]int)
	ansArr := []int{}
	maxCount := 0
	var postorderTraversal func(root *TreeNode) int
	postorderTraversal = func(root *TreeNode) int {
		ret := 0
		if root.Left != nil {
			ret += postorderTraversal(root.Left)
		}
		if root.Right != nil {
			ret += postorderTraversal(root.Right)
		}
		//计算当前节点的子树元素和
		ans := root.Val + ret
		// ansArr = append(ansArr, ans)
		total[ans]++
		if total[ans] > maxCount {
			maxCount = total[ans]
		}
		return ans
	}
	postorderTraversal(root)
	for key, value := range total {
		if value == maxCount {
			ansArr = append(ansArr, key)
		}
	}
	return ansArr
}
