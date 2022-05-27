package tree

func isUnivalTree(root *TreeNode) bool {
	var dfs func(rootValue int, root *TreeNode) bool
	dfs = func(rootValue int, root *TreeNode) bool {
		if root.Val != rootValue {
			return false
		}
		if root.Left != nil {
			result := dfs(rootValue, root.Left)
			if !result {
				return result
			}
		}
		if root.Right != nil {
			result := dfs(rootValue, root.Right)
			if !result {
				return result
			}
		}
		return true
	}
	return dfs(root.Val, root)

}
