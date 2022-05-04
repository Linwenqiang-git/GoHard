package tree

import (
	utility "github.linwenqiang.com/LeetCode/Utility"
)

func levelOrderBottom(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	deep := getMaxTreeDeep(root)
	for i := 0; i < deep; i++ {
		ret = append(ret, []int{})
	}
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {

		if ret[deep-1-level] == nil {
			ret[deep-1-level] = []int{}
		}
		p := []*TreeNode{}
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			ret[deep-1-level] = append(ret[deep-1-level], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		queue = p
	}
	return ret
}

func getMaxTreeDeep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return utility.Max(getMaxTreeDeep(root.Left), getMaxTreeDeep(root.Right)) + 1
}
