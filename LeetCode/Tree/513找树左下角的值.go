package tree

func findBottomLeftValue(root *TreeNode) int {
	ret := levelOrder_513(root)
	return ret[len(ret)-1][0]
}
func levelOrder_513(root *TreeNode) [][]int {
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
