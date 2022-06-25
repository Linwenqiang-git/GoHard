package dynamicprogramming

import (
	utility "github.linwenqiang.com/LeetCode/Utility"
)

func minCost(costs [][]int) int {
	n := len(costs)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 3)
		if i == 0 {
			dp[i] = costs[i]
			continue
		}
		for j := 0; j < 3; j++ {
			dp[i][j] = utility.Min(dp[i-1][(j+1)%3], dp[i-1][(j+2)%3]) + costs[i][j]
		}
	}
	minCost := 1 << 31
	for _, value := range dp[n-1] {
		if value < minCost {
			minCost = value
		}
	}
	return minCost
}
