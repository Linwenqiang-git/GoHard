package tree

func numTrees(n int) int {
	return numTree(1, n)
}

func numTree(start, end int) int {
	if start > end {
		return 1
	}
	result := 0
	for i := start; i <= end; i++ {
		leftTreeNodes := numTree(start, i-1)
		rightTreeNodes := numTree(i+1, end)
		//分别从左侧和右侧选一个,排列组合
		result += leftTreeNodes * rightTreeNodes
	}
	return result
}

//递归的方式会导致超时，且对于同一个节点最为根，numTree会被反复调用
//可以用数学的方式推倒
/*
以 i 为根、长度为 n 的不同二叉搜索树 得到F(i,n)
G(i−1) 表示构建不同左子树的数量
G(n-i) 表示构建不同右子树的数量
则 F(i,n) = G(i−1) * G(n-i)
则总数 从 i= 1开始到i = n的累加
*/
func numTrees2(n int) int {
	// 当i=0、i=1时，空数和一个节点时，都只有一种方案
	//G[n] = G(0) * G(n-1) + G(1) * G(n-2) + ... + G(n-1) * G(0)
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}
