package tree

import (
	utility "github.linwenqiang.com/LeetCode/Utility"
)

func minDepth(root *TreeNode) int {
	return deap(root)
}

func deap(root *TreeNode) int {
	deep := 1
	if root.Left == nil || root.Right == nil {
		return deep
	} else {
		leftHeight := 0
		rightHeight := 0
		leftHeight += minDepth(root.Left) + 1
		rightHeight += minDepth(root.Right) + 1
		return utility.Min(leftHeight, rightHeight)
	}
}
