package dynamicprogramming

/*
题目：你现在手里有一份大小为 n x n 的 网格 grid，上面的每个 单元格 都用 0 和 1 标记好了。其中 0 代表海洋，1 代表陆地。
请你找出一个海洋单元格，这个海洋单元格到离它最近的陆地单元格的距离是最大的，并返回该距离。如果网格上只有陆地或者海洋，请返回 -1。
我们这里说的距离是「曼哈顿距离」（ Manhattan Distance）：(x0, y0) 和 (x1, y1) 这两个单元格之间的距离是 |x0 - x1| + |y0 - y1| 。

题解：点(i,j)若是海洋，则距离他最近的陆地可能来自上下左右四个方向
f(i,j)表示其距离最近陆地的曼哈顿距离，
对于左边和上边而言
f(i,j) = {
	min(f(i-1,j),f(i,j-1)) + 1  u(i,j) = 0  取二者的最小值是因为要取最近的陆地
	0							u(i,j) = 1
}
对于右边和下边而言
f(i,j) = {
	min(f(i+1,j),f(i,j+1)) + 1  u(i,j) = 0  取二者的最小值是因为要取最近的陆地
	0							u(i,j) = 1
}
*/

func maxDistance(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	data := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				data[i][j] = 0
			} else {
				data[i][j] = 100 * 100
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				data[i][j] = 0
			} else {

			}
		}
	}
	return 0
}
