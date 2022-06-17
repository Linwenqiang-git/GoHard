package algorithm

func duplicateZeros(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			if i+1 < n {
				for j := n - 1; j > i+1; j-- {
					arr[j] = arr[j-1]
				}
				i += 1
				arr[i] = 0
			}
		}
	}
}
