package tree

//关键点：已知左子树和右子树的最大深度l和r，那么二叉树的最大深度为max(l,r)+1
//同理 l和r的最大深度也可以采用该方式获取

//深度优先算法
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
