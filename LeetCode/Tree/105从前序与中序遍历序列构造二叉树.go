package tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	root := &TreeNode{}
	root.Val = preorder[0]//先序遍历的第一个为根节点
	temp := root

	return root
}
