package tree

/*
数组是严格递增
二叉搜索树，中序遍历的结果就是递增的，可以按照中序的方式生成树
*/
func sortedArrayToBST(nums []int) *TreeNode {
	var root = &TreeNode{}
	middle := len(nums) / 2
	if len(nums) >= 1 {
		root.Val = nums[middle]
		root.Left = sortedArrayToBST(nums[:middle])
		root.Right = sortedArrayToBST(nums[middle+1:])
		return root
	} else {
		return nil
	}

}
