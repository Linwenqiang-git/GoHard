package interviewquestion

import (
	. "github.linwenqiang.com/LeetCode/Tree"
)

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	queue := []*TreeNode{root}
	var pre, cur *TreeNode = nil, root
	for len(queue) > 0 || cur != nil {
		for cur != nil {
			queue = append(queue, cur)
			cur = cur.Left
		}
		cur = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if pre == p {
			return cur
		}
		pre = cur
		cur = cur.Right
	}
	return nil
}

/*
二叉搜索树的一个性质是中序遍历序列单调递增，因此二叉搜索树中的节点 ppp 的后继节点满足以下条件：

    后继节点的节点值大于 p 的节点值；

    后继节点是节点值大于 p 的节点值的所有节点中节点值最小的一个节点。
*/
func inorderSuccessor2(root *TreeNode, p *TreeNode) *TreeNode {
	var successor *TreeNode
	//存在右侧节点
	if p.Right != nil {
		successor = p.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		return successor
	}
	node := root
	for node != nil {
		//当值比p大，在向左侧找
		if node.Val > p.Val {
			successor = node
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return successor
}
