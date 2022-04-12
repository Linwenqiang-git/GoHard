package matrix

import (
	utility "github.linwenqiang.com/LeetCode/Utility"
)

func imageSmoother(img [][]int) [][]int {
	row := len(img)
	column := len(img[0])
	result := make([][]int, row)
	for row_index := range result {
		result[row_index] = make([]int, column)
	}

	for row_index := 0; row_index < row; row_index++ {
		for column_index := 0; column_index < column; column_index++ {
			//从上面开始，按着顺时针相加
			sum := 0
			num := 0
			//上
			if row_index-1 >= 0 {
				sum += img[row_index-1][column_index]
				num++
			}
			//右上
			if row_index-1 >= 0 && column_index+1 < column {
				sum += img[row_index-1][column_index+1]
				num++
			}
			//右
			if column_index+1 < column {
				sum += img[row_index][column_index+1]
				num++
			}
			//右下
			if row_index+1 < row && column_index+1 < column {
				sum += img[row_index+1][column_index+1]
				num++
			}
			//下
			if row_index+1 < row {
				sum += img[row_index+1][column_index]
				num++
			}
			//左下
			if row_index+1 < row && column_index-1 >= 0 {
				sum += img[row_index+1][column_index-1]
				num++
			}
			//左
			if column_index-1 >= 0 {
				sum += img[row_index][column_index-1]
				num++
			}
			//左上
			if row_index-1 >= 0 && column_index-1 >= 0 {
				sum += img[row_index-1][column_index-1]
				num++
			}
			sum += img[row_index][column_index]
			num++
			result[row_index][column_index] = sum / num
		}
	}
	return result
}

//官方解法，用额外的两个循环代替上面的if
func imageSmoother2(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			sum, num := 0, 0
			for _, row := range img[utility.Max(i-1, 0):utility.Min(i+2, m)] {
				for _, v := range row[utility.Max(j-1, 0):utility.Min(j+2, n)] {
					sum += v
					num++
				}
			}
			ans[i][j] = sum / num
		}
	}
	return ans
}
