package matrix

/*
原地旋转方案
难点：是要旋转哪些位置
规律 （x,y）=>(x1,y1)
	x1 = y
	y1 = n-x
	n = len(matrix) - 1
*/

/*
用翻转代替旋转
比较灵活，代码易理解
先水平对折，在对角线对折
*/
func rotate(matrix [][]int) {
	n := len(matrix)
	//水平翻转
	for i := 0; i < n/2; i++ {
		temp := matrix[i]
		matrix[i] = matrix[n-i-1]
		matrix[n-i-1] = temp
		//写法2
		//matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	//对折翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
			//写法2
			//matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
