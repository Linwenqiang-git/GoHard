package matrix

import (
	utility "github.linwenqiang.com/LeetCode/Utility"
)

/*
动态规划的解题思路
动态转移方程如下：
	对于第一行：
	i=0,j>0  dp[0][j] = dp[0][j-1] + grid[0][j]
	对于第一列：
	i>0,j=0  dp[i][0] = dp[i-1][0] + grid[i][0]
	对于表格中间的数据：要么从上一个格子下移，要嘛做左边的格子右移
	i>0,j>0  dp[i][j] = min(dp[i-1][j],dp[i][j-1])+grid[i][j])
*/
func minPathSum(grid [][]int) int {
	row := len(grid)
	column := len(grid[0])
	if row == 0 || column == 0 {
		return 0
	}
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, column)
	}
	dp[0][0] = grid[0][0]
	for x := 0; x < row; x++ {
		for y := 0; y < column; y++ {
			if x == 0 && y > 0 {
				dp[0][y] = dp[0][y-1] + grid[0][y]
			} else if y == 0 && x > 0 {
				dp[x][0] = dp[x-1][0] + grid[x][0]
			} else if x > 0 && y > 0 {
				dp[x][y] = utility.Min(dp[x-1][y], dp[x][y-1]) + grid[x][y]
			}
		}
	}
	return dp[row-1][column-1]
}
