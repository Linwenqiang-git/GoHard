package tree

import "fmt"

func isSameTree(p *TreeNode, q *TreeNode) bool {
	leftTree := getTreeNode(p)
	rightTree := getTreeNode(q)
	fmt.Printf("%v\n", leftTree)
	fmt.Printf("%v\n", rightTree)
	if len(leftTree) != len(rightTree) {
		return false
	}
	for i := 0; i < len(leftTree); i++ {
		if leftTree[i] != rightTree[i] {
			return false
		}
	}
	return true
}

func getTreeNode(p *TreeNode) []int {
	result := []int{}
	if p == nil {
		return nil
	}
	result = append(result, p.Val)
	if p.Left != nil {
		result = append(result, getTreeNode(p.Left)...)
	}

	if p.Right != nil {
		if p.Left == nil {
			result = append(result, 0)
		}
		result = append(result, getTreeNode(p.Right)...)
	}
	return result
}

func main() {

}
