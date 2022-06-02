package tree

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root != nil && root.Val == key {
		accessor := root.Right
		for accessor.Left != nil {
			accessor = accessor.Left
		}
		value := accessor.Val
		accessor.Right = firstTraversal2(root.Right, value)
		accessor.Left = root.Left
		return accessor
	} else {
		return firstTraversal2(root, key)
	}
}
func firstTraversal2(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Right != nil {
			if root.Right.Left != nil {
				access := root.Right.Left
				for access.Left != nil{
					access = access.Left
				}
				root.Val = access.Val
				

			} else {
				left := root.Left
				root = root.Right
				root.Left = left
			}

		} else if root.Left != nil {
			root = root.Left
			root.Right = nil
		} else {
			root = nil
		}

	} else if root.Val > key {
		root.Left = firstTraversal2(root.Left, key)
	} else {
		root.Right = firstTraversal2(root.Right, key)
	}
	return root
}

func firstTraversal(root *TreeNode, key int) []int {
	result := []int{}
	if root.Val != key {
		result = append(result, root.Val)
	}
	if root.Left != nil {
		result = append(result, firstTraversal(root.Left, key)...)
	}
	if root.Right != nil {
		result = append(result, firstTraversal(root.Right, key)...)
	}
	return result
}
