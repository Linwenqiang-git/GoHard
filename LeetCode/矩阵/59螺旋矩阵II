package matrix

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}
	//移动方向，按照右->下->左->上
	direct := [][]int{
		{0, 1},  //右
		{1, 0},  //下
		{0, -1}, //左
		{-1, 0}} //上
	step := 1
	direct_index := 0
	i, j := 0, 0
	for step <= n*n {
		if result[i][j] == 0 {
			result[i][j] = step
			step++
		}
		if i+direct[direct_index][0] >= n || //碰触到下边界
			i+direct[direct_index][0] < 0 || //碰触到上边界
			j+direct[direct_index][1] >= n || //碰触到右边界
			j+direct[direct_index][1] < 0 || //碰触到最边界
			result[i+direct[direct_index][0]][j+direct[direct_index][1]] != 0 { //下个方向是已经访问的表格
			//转换方向
			direct_index++
			direct_index = direct_index % 4
		} else {
			i += direct[direct_index][0]
			j += direct[direct_index][1]
		}
	}
	return result
}
