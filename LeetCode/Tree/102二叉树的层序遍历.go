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

//广度优先
func levelOrder2(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	//辅助队列
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		if ret[i] == nil {
			ret[i] = make([]int, 0)
		}
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}
