package dynamicprogramming

//对于到达点 f(i,j) 只能由f(i-1,j)右移 或者 f(i,j-1) 下移得到
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	f := make([]int, m)
	if obstacleGrid[0][0] == 0 {
		f[0] = 1
	}
	//两层循环，第二层在第一层的基础上做操作
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
				continue
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				//当到达第二层之后，就实现了下移和右移
				f[j] += f[j-1]
			}
		}
	}
	//滚动数组的最后一个点，即为结束点
	return f[len(f)-1]
}
