package chart

func findCircleNum(isConnected [][]int) int {
	related := make([][]int, len(isConnected))
	visited := make(map[int]bool)
	for i := 0; i < len(related); i++ {
		related[i] = []int{}
	}
	for i, city := range isConnected {
		for j, info := range city {
			if i == j {
				continue
			}
			if info == 1 {
				related[i] = append(related[i], j)
			}
		}
	}
	var dfs func(i int)
	dfs = func(i int) {
		visited[i] = true
		for _, value := range related[i] {
			if _, ok := visited[value]; !ok {
				dfs(value)
			}
		}
	}

	province := 0
	for i := 0; i < len(related); i++ {
		if _, ok := visited[i]; !ok {
			dfs(i)
			province++
		}
	}
	return province
}
