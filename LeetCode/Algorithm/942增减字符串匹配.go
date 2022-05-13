package algorithm

/*
当s[0]为I时，perm[0] = 0 可以保证任意的perm[1] > perm[0]
当s[0]为D时，perm[0] = n 可以保证任意的perm[1] < perm[0]
同理，perm剩下的部分也按照这个思路，
当为I时，取当前剩下的最小值，当为D时，取当前剩下的最大值
*/
func diStringMatch(s string) []int {
	n := len(s)
	ret := []int{}
	if n == 0 {
		return ret
	}
	low, height := 0, n
	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			ret = append(ret, low)
			low++
		} else {
			ret = append(ret, height)
			height--
		}
	}
	ret = append(ret, low) //最后剩下一个数，此时 lo == hi
	return ret
}
