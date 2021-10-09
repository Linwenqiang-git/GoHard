package tree

/*
二叉查找树（Binary Search Tree），（又：二叉搜索树，二叉排序树）它或者是一棵空树，
或者是具有下列性质的二叉树：
若它的左子树不空，则左子树上所有结点的值均 小于 它的根结点的值；
若它的右子树不空，则右子树上所有结点的值均 大于 它的根结点的值；
它的左、右子树也分别为二叉排序树

给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
*/

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateTree(1, n)
}

func generateTree(start, end int) []*TreeNode {
	result := []*TreeNode{}
	if start > end {
		return nil
	}
	for i := start; i <= end; i++ {
		currentPoint := &TreeNode{Val: i, Left: nil, Right: nil}
		leftTreeNode := generateTree(1, i-1)
		rightTreeNode := generateTree(i+1, end)
		//从左侧选一个 在从右侧选一个
		if leftTreeNode == nil && rightTreeNode == nil {
			result = append(result, currentPoint)
			continue
		}
		if leftTreeNode == nil && rightTreeNode != nil {
			for r := 0; r < len(rightTreeNode); r++ {
				currentPoint.Right = rightTreeNode[r]
				result = append(result, currentPoint)
			}
			continue
		}
		if leftTreeNode != nil && rightTreeNode == nil {
			for l := 0; l < len(leftTreeNode); l++ {
				currentPoint.Right = leftTreeNode[l]
				result = append(result, currentPoint)
			}
			continue
		}
		for l := 0; l < len(leftTreeNode); l++ {
			currentPoint.Left = leftTreeNode[l]
			for r := 0; r < len(rightTreeNode); r++ {
				currentPoint.Right = rightTreeNode[r]
				result = append(result, currentPoint)
			}
		}
	}
	return result
}

func main() {

}
