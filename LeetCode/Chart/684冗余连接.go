package chart

func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1) //这里+1是因为节点从1开始
	//设置每个节点的上级是自己，即初始时，每个元素都是一个连通分量
	for i, _ := range parent {
		parent[i] = i
	}
	//返回节点x的集合代表
	var find func(x int) int
	find = func(x int) int {
		//上级不是自己，则一层层的向上找
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	var join func(from, to int) bool
	join = func(from, to int) bool {
		x, y := find(from), find(to)
		//from和to的代表节点在这之前是联通的，当前边会造成环
		if x == y {
			return false
		}
		//from和to在遍历到这里之前还是不连通的，所以不会存在环，将二者连通
		parent[x] = y
		return true
	}

	for _, val := range edges {
		if !join(val[0], val[1]) {
			return val
		}
	}
	return nil
}
