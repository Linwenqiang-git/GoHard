package algorithm

import "sort"

func makesquare(matchsticks []int) bool {
	totalLen := 0
	for _, l := range matchsticks {
		totalLen += l
	}
	if totalLen%4 != 0 {
		return false
	}
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks))) // 减少搜索量

	edges := [4]int{}
	/*
		思路是:
		先取第一根放在第一个边上，如果不适合，在放第二条边，依次放完四条边，如果都不合适，
		那就返回false，如果中间哪个边合适了，就继续取第二根

		第二根放的思路和第一根一样，都是依次从第一条边开始试
		直到最后一根也都放进去了，则能拼成一个正方形
	*/
	var dfs func(int) bool
	dfs = func(idx int) bool {
		if idx == len(matchsticks) {
			return true
		}
		for i := range edges {
			edges[i] += matchsticks[idx]
			if edges[i] <= totalLen/4 && dfs(idx+1) {
				return true
			}
			edges[i] -= matchsticks[idx]
		}
		return false
	}
	return dfs(0)
}

func Call_473() {
	makesquare([]int{1, 1, 2, 2, 2})
}
