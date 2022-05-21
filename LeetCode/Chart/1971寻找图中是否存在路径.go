package chart

func validPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination{
        return true
    }
    if len(edges) == 0{
        return false
    }
	//初始化一个dig,存储对于当前节点n,可到达的节点有哪些
	digs := make([][]int, n)
	for i := 0; i < n; i++ {
		digs[i] = []int{}
	}
	for i := 0; i < len(edges); i++ {
		digs[edges[i][0]] = append(digs[edges[i][0]], edges[i][1])
		digs[edges[i][1]] = append(digs[edges[i][1]], edges[i][0])
	}
	//获取源头可到达的节点
	queue := digs[source]
	acrossMap := make(map[int]bool)
	for i := 0; i < n; i++ {
		acrossMap[i] = false
	}
	acrossMap[source] = true
	for len(queue) > 0 {
		node := queue[0]
		if !acrossMap[node] {
			acrossMap[node] = true
			queue = append(queue, digs[node]...)
		}

		if node == destination {
			return true
		}
		queue = queue[1:]	
	}
	return false
}
