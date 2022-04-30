package tree

var result [][]int

func levelOrder(root *TreeNode) [][]int {
	result = make([][]int, 0)
	ergodic(root, 0)
	return result
}

func ergodic(root *TreeNode, index int) {
	if root != nil {
		if index >= len(result) {
			result = append(result, make([]int, 0))
		}
		if result[index] == nil {
			result[index] = make([]int, 0)
		}
		result[index] = append(result[index], root.Val)
		ergodic(root.Left, index+1)
		ergodic(root.Right, index+1)
	}
}
