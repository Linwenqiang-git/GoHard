package tree

import "github.com/go-playground/locales/root"

var result = make([][]int, 0)

func levelOrder(root *TreeNode) [][]int {

}

func ergodic(left *TreeNode, right *TreeNode) []int{
	if left == right {
		result = append(result, []int{left.Val})
		return
	}
	result = append(result, ergodic(left.Left,left.Right))
}
