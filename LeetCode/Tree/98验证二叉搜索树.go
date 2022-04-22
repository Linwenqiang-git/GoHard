package tree

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	values := postorderTraversal(root)
	for i := 0; i < len(values)-1; i++ {
		if values[i] >= values[i+1] {
			return false
		}
	}
	return true
}

//中序遍历得到的结果是升序的数组
func postorderTraversal(node *TreeNode) []int {

	result := make([]int, 0)
	if node.Left != nil {
		result = append(result, postorderTraversal(node.Left)...)
	}
	result = append(result, node.Val)
	if node.Right != nil {
		result = append(result, postorderTraversal(node.Right)...)
	}
	return result
}
