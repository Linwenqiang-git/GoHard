package tree

func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	//辅助队列
	q := []*TreeNode{root}
	for level := 0; len(q) > 0; level++ {
		ret = append(ret, []int{})
		if ret[level] == nil {
			ret[level] = make([]int, 0)
		}
		p := []*TreeNode{}
		data := []int{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			data = append(data, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		if level%2 != 0 {
			//从右往左
			for i, n := 0, len(data); i < n/2; i++ {
				data[i], data[n-1-i] = data[n-1-i], data[i]
			}
		}
		ret[level] = append(ret[level], data...)
		q = p
	}
	return ret
}
