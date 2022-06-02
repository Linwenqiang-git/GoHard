package matrix

import (
	"fmt"
	"strconv"
)

func maxAreaOfIsland(grid [][]int) int {
	var dir = [][]int{{-1, 0}, {0, -1}} //上和左
	width := len(grid)
	height := len(grid[0])

	parent := make(map[string]string)
	areaCount := make(map[string]int)

	var find func(node string) string
	find = func(node string) string {
		if parent[node] != node {
			parent[node] = find(parent[node])
		}
		return parent[node]
	}
	var join func(node1, node2 string)
	join = func(nodeParent, topNodeParent string) {
		parent[topNodeParent] = nodeParent
	}

	for i, column := range grid {
		for j, value := range column {
			if value == 1 {
				key := strconv.Itoa(i) + strconv.Itoa(j)
				//默认上级节点是自己
				parent[key] = key
				//给当前节点找一个最近的上级节点
				for _, d := range dir {
					x := i + d[0]
					y := j + d[1]
					if x >= 0 && y >= 0 && x < width && y < height && grid[x][y] == 1 {
						//左移时已经找到相邻节点
						if parent[key] != key {
							topKey := strconv.Itoa(x) + strconv.Itoa(y)
							topKeyParent := find(topKey)
							join(parent[key], topKeyParent)
							break
						}
						parentKey := strconv.Itoa(x) + strconv.Itoa(y)
						parent[key] = find(parentKey)
					}
				}

			}
		}
	}
	for _, v := range parent {
		areaCount[v]++
	}
	maxarea := 0
	for k, v := range areaCount {
		fmt.Println(k, v)
		if maxarea < v {
			maxarea = v
		}
	}
	return maxarea
}
