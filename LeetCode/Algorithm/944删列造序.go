package algorithm

func minDeletionSize(strs []string) (deleteCount int) {
	n := len(strs)
	deleteCount = 0
	if n == 0 {
		return deleteCount
	}
	c := len(strs[0])
	for i := 0; i < c; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				if strs[j][i] < strs[j-1][i] {
					deleteCount++
					break
				}
			}

		}
	}
	return
}
