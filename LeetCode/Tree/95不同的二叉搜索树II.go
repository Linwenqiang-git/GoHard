package tree

import "fmt"

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
	fmt.Printf("start = %d", start)
	if start > end {
		return []*TreeNode{nil} //这里不要返回nil 这样下面遍历就不需要考虑多种情况
	}
	result := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftTreeNode := generateTree(start, i-1) //左侧要传start 不能传1，比如当1是根节点的时候，子节点就不能有1
		rightTreeNode := generateTree(i+1, end)
		//从左侧选一个 在从右侧选一个
		for _, left := range leftTreeNode {
			for _, right := range rightTreeNode {
				currentPoint := &TreeNode{Val: i, Left: nil, Right: nil}
				currentPoint.Left = left
				currentPoint.Right = right
				result = append(result, currentPoint)
			}
		}
	}
	return result
}

func main() {

}
