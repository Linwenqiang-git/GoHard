package algorithm

func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	for  num:= 0;num<k;num++{
		for i := m-1;i>=0;i--{
			for j:=n-1;j>=0;j--{
				grid[i][j] = 1
			}
		}
	}
	return grid
}