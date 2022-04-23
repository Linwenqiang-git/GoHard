package tree

import (
	"math"

	utility "github.linwenqiang.com/LeetCode/Utility"
)

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//深度优先，找到叶子节点
	if root.Left == nil && root.Right == nil {
		return 1
	}
	min := math.MaxInt32
	if root.Left != nil {
		min = utility.Min(minDepth(root.Left), min)
	}
	if root.Right != nil {
		min = utility.Min(minDepth(root.Right), min)
	}
	return min + 1 //+1是加上自己本身
}
