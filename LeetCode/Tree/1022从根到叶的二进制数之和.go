package tree

import (
	"fmt"
	"strconv"

	utility "github.linwenqiang.com/LeetCode/Utility"
)

var sum int

func sumRootToLeaf(root *TreeNode) int {
	sum = 0
	dfs(root, "")
	return sum
}
func dfs(root *TreeNode, s string) string {
	s += strconv.Itoa(root.Val)
	if root.Left != nil {
		dfs(root.Left, s)
	}
	if root.Right != nil {
		dfs(root.Right, s)
	}
	if root.Left == nil && root.Right == nil {
		fmt.Println(s)
		sum += int(utility.BinaryToDecimal(s))
	}
	s = s[:len(s)-1]
	return s
}
