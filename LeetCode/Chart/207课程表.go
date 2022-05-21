package chart

func canFinish(numCourses int, prerequisites [][]int) bool {
	node_edge_nodes := make([][]int, numCourses)
	//栈结构，存储节点是否被访问过
	visitied := make([]int, numCourses)
	//存储一种可能的拓扑排序（本题不需要返回拓扑结构，可不使用这个变量）
	topology := []int{}
	//当前遍历的节点是否还有效
	valid := true
	//深度搜索方法声明
	var dfs func(nodeNum int)

	dfs = func(nodeNum int) {
		visitied[nodeNum] = 1 //1表示搜索中
		for _, node := range node_edge_nodes[nodeNum] {
			if visitied[node] == 0 {
				dfs(node)
				if !valid {
					return
				}
			} else if visitied[node] == 1 { //如果当前节点已经在搜索中，且又在某个节点的可达点，则不符合本题的题意
				valid = false
				return
			}
		}
		visitied[nodeNum] = 2 //2表示搜索完成
		topology = append(topology, nodeNum)
	}
	//先找到每个节点对应可到达的边
	for _, info := range prerequisites {
		node_edge_nodes[info[1]] = append(node_edge_nodes[info[1]], info[0]) //由1到0
	}
	//深度搜索每个节点
	for i := 0; i < numCourses && valid; i++ {
		if visitied[i] == 0 { //0表示未搜索
			dfs(i)
		}
	}
	return valid
}
