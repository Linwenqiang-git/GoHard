package matrix

/*
关键：
1.对已经访问的点的记录
2.关键节点的位置变换
*/
func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	column := len(matrix[0])
	var result = make([]int, 0)
	if row == 0 || column == 0 {
		return result
	}
	//记录已访问的矩阵
	visit_matrix := make([][]int, row)
	for i := range matrix {
		visit_matrix[i] = make([]int, column)
	}
	//路径长度，结束标记
	rootLength := 0
	//移动方向，按照右->下->左->上
	direct := [][]int{
		{0, 1},  //右
		{1, 0},  //下
		{0, -1}, //左
		{-1, 0}} //上
	direct_index := 0
	i, j := 0, 0
	for rootLength < row*column {
		if visit_matrix[i][j] != 1 {
			result = append(result, matrix[i][j])
			visit_matrix[i][j] = 1
			rootLength++
		}
		if i+direct[direct_index][0] >= row || //碰触到下边界
			i+direct[direct_index][0] < 0 || //碰触到上边界
			j+direct[direct_index][1] >= column || //碰触到右边界
			j+direct[direct_index][1] < 0 || //碰触到最边界
			visit_matrix[i+direct[direct_index][0]][j+direct[direct_index][1]] == 1 { //下个方向是已经访问的表格
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
