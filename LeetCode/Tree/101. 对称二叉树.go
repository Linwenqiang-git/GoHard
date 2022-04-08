package tree

//关键点：每一棵树的左子树都与另一个树的左子树镜像对称
//递归的方式实现
func check(p, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}
	if q == nil || p == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}
