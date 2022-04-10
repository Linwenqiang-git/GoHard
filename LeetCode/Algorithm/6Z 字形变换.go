package algorithm

import (
	"fmt"
)

func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 || numRows >= n {
		return s
	}
	dip := make([][]byte, numRows)
	//每填充一个周期需要的个数
	t := 2*numRows - 2
	//每个周期占用的列数
	r := numRows - 1
	//一共有多少个周期(最后一个视为完整周期,如果n刚好是完整周期，则-1在加一个周期是对的，不足下个周期，-1也不影响)
	c := (n - 1 + t) / t
	//一共有多少列
	column_nums := c * r
	for i := 0; i < numRows; i++ {
		dip[i] = make([]byte, column_nums)
	}
	column := 0
	row := 0
	dir := 1 //1向下填充 2右上填充
	for p := 0; p < n; p++ {
		fmt.Printf("row = %d,column = %d\n", row, column)
		dip[row][column] = s[p]
		if dir == 1 {
			row++
			if row == numRows {
				row -= 2
				//切换到下一列
				column++
				if row > 0 {
					dir = 2
				}
			}
		} else {
			row--
			if row <= 0 {
				dir = 1
			}
			//切换到下一列
			column++
		}
	}
	result := ""
	for _, row := range dip {
		for _, ch := range row {
			if ch > 0 {
				result += string(ch)
			}
		}
	}
	return result
}

func Call_6() {
	fmt.Println(convert("ABXD", 2))
}
