package dynamicprogramming

func lenLongestFibSubseq(arr []int) int {
	ans := 0
	n := len(arr)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	dic := make(map[int]int)
	for i, v := range arr {
		dic[v] = i
	}
	for i, v := range arr {
		for j := n - 1; j >= 0 && arr[j]*2 > v; j-- {
			if k, ok := dic[v-arr[j]]; ok {
				dp[j][i] = max(dp[k][j]+1, 3)
				ans = max(ans, dp[j][i])
			}
		}
	}
	return ans
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
