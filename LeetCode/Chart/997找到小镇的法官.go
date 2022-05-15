package chart

/*
可看成一个有向图，每个人都是图上的一个节点
trust[i] 对应的数值，就是节点到节点之间是否可达
*/
func findJudge(n int, trust [][]int) int {
	//存放每个节点的入和出
	//makeIn[1] 存放第一个人的入度
	//makeOut[1]存放第一个人的出度
	makeIn := make([]int, n+1)
	makeOut := make([]int, n+1)
	for i := 0; i < len(trust); i++ {
		makeIn[trust[i][1]]++
		makeOut[trust[i][0]]++
	}
	for i := 1; i < n+1; i++ {
		if makeIn[i] == n-1 && makeOut[i] == 0 {
			return i
		}
	}
	return -1
}
