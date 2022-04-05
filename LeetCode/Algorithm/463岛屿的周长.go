package main

import "fmt"

func main() {
	//data := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	data := [][]int{{1, 0}}
	fmt.Println(islandPerimeter(data))
}

type pair struct {
	x int
	y int
}

//格子移动的方向 上、下、右、左
var dir = []pair{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

//迭代算法
func islandPerimeter(grid [][]int) int {
	ans := 0
	x_len := len(grid[0]) //水平方向最大值
	y_len := len(grid)    //垂直方向的最大值
	for row_index, row := range grid {
		for column_index, value := range row {
			if value == 1 {
				for _, dir_value := range dir {
					x := dir_value.x + row_index
					y := dir_value.y + column_index
					if x < 0 || //上移超过边界
						x >= y_len || //下移超过右边界
						y < 0 || //左移超过上边界
						y >= x_len || //右移超过下边界
						grid[x][y] == 0 { //移动到水中
						ans += 1
					}
				}
			}
		}
	}
	return ans
}
