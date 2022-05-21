package chart

//广度搜索，一层一层的遍历，需要额外的辅助空间记录已经访问的点，防止陷入死循环
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	queue := []*Node{node}
	visit := map[*Node]*Node{}
	visit[node] = &Node{node.Val, []*Node{}}
	for len(queue) > 0 {
		_childnodes := queue[0]
		queue = queue[1:]
		for _, neighbor := range _childnodes.Neighbors {
			if _, ok := visit[neighbor]; !ok {
				//没被访问，就添加一下
				visit[neighbor] = &Node{neighbor.Val, []*Node{}}
				//将邻居节点加入队列
				queue = append(queue, neighbor)
			}
			//更新当前节点的邻居节点
			visit[_childnodes].Neighbors = append(visit[_childnodes].Neighbors, visit[neighbor])
		}
	}
	return visit[node]
}
