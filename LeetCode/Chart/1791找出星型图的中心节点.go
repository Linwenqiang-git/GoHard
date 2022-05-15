package chart

func findCenter(edges [][]int) int {
	n := len(edges) + 1
	line := make([]int, n+1)
	for i := 0; i < len(edges); i++ {
		line[edges[i][0]]++
		line[edges[i][1]]++
	}
	ret := 0
	for index, data := range line {
		if data == len(edges) {
			ret = index
			break
		}
	}
	return ret
}
