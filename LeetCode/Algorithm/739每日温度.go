package algorithm

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ret := []int{}
	if n == 0 {
		return ret
	}
	for i := 0; i < n; i++ {
		tem := temperatures[i]
		days := 0
		day := 0
		for j := i + 1; j < n; j++ {
			day++
			if temperatures[j] > tem {
				days = day
				break
			}
		}
		ret = append(ret, days)
	}
	return ret
}
