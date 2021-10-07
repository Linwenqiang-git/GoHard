//Definition for a binary tree node.
package tree

//二叉树的中序遍历就是首先遍历左子树，然后访问当前节点，最后遍历右子树。

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		if root.Left != nil {
			result = append(result, inorderTraversal(root.Left)...)
		}
		result = append(result, root.Val)
		if root.Right != nil {
			result = append(result, inorderTraversal(root.Right)...)
		}
	}
	return result
}

func main() {

}
