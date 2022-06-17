package matrix

func findDiagonalOrder(mat [][]int) []int {
	dir_weight_1 := [][]int{{-1, 1}, {0, 1}, {1, 0}} //左下到右上、右、下
	dir_weight_2 := [][]int{{1, -1}, {1, 0}, {0, 1}} //右上到左下、下、右
	n := len(mat)
	m := len(mat[0])
	dir := 1 //1左下到右上 2右上到左下
	ans := []int{}
	i, j := 0, 0
	for len(ans) < m*n {
		ans = append(ans, mat[i][j])
		if dir == 1 {
			for index, val := range dir_weight_1 {
				if i+val[0] >= 0 && i+val[0] < n && j+val[1] >= 0 && j+val[1] < m {
					if index > 0 {
						dir = 2
					}
					i += val[0]
					j += val[1]
					break
				} 
			}
		} else {
			for index, val := range dir_weight_2 {
				if i+val[0] >= 0 && i+val[0] < n && j+val[1] >= 0 && j+val[1] < m {
					if index > 0 {
						dir = 1
					}
					i += val[0]
						j += val[1]
						break
				} 
			}
		}
	}
	return ans
}
