package chain

func insert(aNode *ListNode, x int) *ListNode {
	node := &ListNode{
		Val:  x,
		Next: nil,
	}
	if aNode == nil {
		node.Next = node
		return node
	}

	headNode := aNode
	isFind := false
	for aNode != headNode {
		if !isFind {
			//位于中间 || 位于最大值 || 位于最小值
			if aNode.Val <= x && x <= aNode.Next.Val {
				node.Next = aNode.Next
				aNode.Next = node
				isFind = true
				break
			}
			if aNode.Next.Val < aNode.Val {
				//当前位于最大节点
				if x >= aNode.Val || x <= aNode.Next.Val {
					node.Next = aNode.Next
					aNode.Next = node
					isFind = true
					break
				}
			}
			if aNode.Next.Val == aNode.Val {
				if aNode == aNode.Next {
					node.Next = aNode.Next
					aNode.Next = node
					isFind = true
					break
				}
			}
		}
		aNode = aNode.Next
	}
	for aNode != headNode {
		aNode = aNode.Next
	}
	if !isFind{
		for aNode.Next !=headNode{
			aNode = aNode.Next
		}
		node.Next = aNode.Next
		aNode.Next = node
		return node.Next
	}

	return aNode
}
