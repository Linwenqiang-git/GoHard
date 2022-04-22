package tree

import (
	"math"
)

/*
题目：给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
	一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。

思路：自顶向下遍历 类似于前序遍历，即 根->左->右
当前节点的高度 = max(左子树，右子树) + 1
会导致height被调用多次
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//因为每个节点都要都要满足，所以还要继续判断根节点的左右节点
	return math.Abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return math.Max(height(root.Left), height(root.Right)) + 1
}

/*
思路2 自底向上 ，类似于后序遍历，即 左->右->根
*/

func isBalanced2(root *TreeNode) bool {
	return height2(root) >= 0
}
func height2(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	leftHeight := height2(root.Left)
	righttHeight := height2(root.Right)
	if math.Abs(leftHeight-righttHeight) > 1 || leftHeight == -1 || righttHeight == -1 {
		return -1
	}
	return math.Max(leftHeight, righttHeight) + 1
}
