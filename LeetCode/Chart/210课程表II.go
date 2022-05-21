package chart

//使用广度搜索实现
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		result         = []int{}
		edge           = make([][]int, numCourses)
		in_degree_nums = make([]int, numCourses)
	)

	for _, info := range prerequisites {
		edge[info[1]] = append(edge[info[1]], info[0]) //添加到当前节点的出度
		in_degree_nums[info[0]]++                      //对于info[0]节点来说 增加一个入度
	}
	queue := []int{}
	//先把没有出度的节点加到队列中，即可以不需要任何前提就可以学习的课程
	for i := 0; i < numCourses; i++ {
		if in_degree_nums[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)
		//对于当前节点下面的出度节点，这些节点的入度都-1
		for _, outNode := range edge[node] {
			in_degree_nums[outNode]--
			if in_degree_nums[outNode] == 0 {
				queue = append(queue, outNode)
			}
		}
	}
	//最后比较一下生成的拓扑节点数量和总数量是否相等
	if len(result) != numCourses {
		return []int{}
	}
	return result
}
