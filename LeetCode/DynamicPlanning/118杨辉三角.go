package dynamicprogramming

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		rowNums := i + 1
		ret[i] = make([]int, rowNums)
		if rowNums == 1 {
			ret[i][0] = 1
		} else {
			for j := 0; j < rowNums; j++ {
				left, right := 0, 0
				if j-1 >= 0 {
					left = ret[i-1][j-1]
				}
				if j < len(ret[i-1]) {
					right = ret[i-1][j]
				}
				ret[i][j] = left + right
			}
		}
	}
	return ret
}
