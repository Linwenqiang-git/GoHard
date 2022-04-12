package matrix

//解题关键：
//对于 x∈[0,mn),第 xxx 个元素在 nums 中对应的下标为(x / n,x % n)
//而在新的重塑矩阵中对应的下标为(x / c,x % c)

func matrixReshape(mat [][]int, r int, c int) [][]int {
	r1 := len(mat)
	c1 := len(mat[0])
	if r1*c1 != r*c {
		return mat
	}
	resharp_mat := make([][]int, r)
	for i := range resharp_mat {
		resharp_mat[i] = make([]int, c)
	}
	for x := 0; x < r1*c1; x++ {
		resharp_mat[x/c][x%c] = mat[x/c1][x%c1]
	}
	return resharp_mat
}

func Call_566() {

}
