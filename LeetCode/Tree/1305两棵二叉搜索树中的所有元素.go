package tree

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	list1 := inorderTraversal(root1)
	list2 := inorderTraversal(root2)
	if len(list1) == 0 {
		return list2
	}
	if len(list2) == 0 {
		return list1
	}
	ret := []int{}
	i, j := 0, 0
	for i < len(list1) || j < len(list2) {
		if i == len(list1) {
			ret = append(ret, list2[j:]...)
			break
		}
		if j == len(list2){
			ret = append(ret, list1[i:]...)
			break
		}
		if list1[i] <= list2[j] {
			ret = append(ret, list1[i])
			i++
		} else {
			ret = append(ret, list2[j])
			j++
		}
	}
	return ret
}

//中序遍历得到升序的数组
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	if root.Left != nil {
		result = append(result, inorderTraversal(root.Left)...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		result = append(result, inorderTraversal(root.Right)...)
	}
	return result
}
