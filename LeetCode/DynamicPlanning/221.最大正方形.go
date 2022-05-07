package dynamicprogramming

import (
	utility "github.linwenqiang.com/LeetCode/Utility"
)

/*
题目：在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

解题思路：对于点(i,j)所能形成的最大面积,即最大边长
 d(i,j) = 1 + min(d(i,j-1),d(i-1,j),d(i-1,j-1))  u(i,j != 0)
 d(i,j) = 0										 u(i,j = 0)

利用滚动数组的思想，每一行记录最大边长
*/

func maximalSquare(matrix [][]byte) int {
	n := len(matrix)
	m := len(matrix[0])
	data := make([][]int, n)
	for i := range data {
		data[i] = make([]int, m)
	}
	maxSide := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '0' {
				data[i][j] = 0
			} else {
				if j == 0 || i == 0 {
					data[i][j] = 1
				} else {
					data[i][j] = 1 + utility.Min(data[i-1][j], utility.Min(data[i][j-1], data[i-1][j-1]))
				}
			}
			if data[i][j] > maxSide {
				maxSide = data[i][j]
			}

		}
	}
	return maxSide * maxSide
}
