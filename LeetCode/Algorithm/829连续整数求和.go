package algorithm

func consecutiveNumbersSum(n int) int {
	ans := 1
	for i := 1; i < n; i++ {
		if size(i, n) {
			ans++
		}
	}
	return ans
}

func size(start, target int) bool {
	for i := start + 1; i <= target; i++ {
		data := (start + i) * (i - start + 1) / 2
		if data == target {
			return true
		}
		if data > target {
			return false
		}
	}
	return false
}
/*
如果 k 是奇数，则当 n 可以被  k 整除时，正整数 n 可以表示成 k 个连续正整数之和；

如果 k 是偶数，则当nn 不可以被 k 整除且 2n 可以被 k 整除时，正整数 n 可以表示成 k 个连续正整数之和。


*/
func isKConsecutive(n, k int) bool {
    if k%2 == 1 {
        return n%k == 0
    }
    return n%k != 0 && 2*n%k == 0
}

func consecutiveNumbersSum2(n int) (ans int) {
    for k := 1; k*(k+1) <= n*2; k++ {
        if isKConsecutive(n, k) {
            ans++
        }
    }
    return
}

