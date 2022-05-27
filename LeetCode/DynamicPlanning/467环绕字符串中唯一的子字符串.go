package dynamicprogramming

import (
	utility "github.linwenqiang.com/LeetCode/Utility"
)

func findSubstringInWraproundString(p string) int {
	//dp[i] 表示p中以字母i结尾且在s中子串的最长长度
	dp := [26]int{}

	k := 0
	/*
	如何计算 dp[i]?
	我们可以在遍历 p 时，维护连续递增的子串长度 k。
	具体来说，遍历到 p[i]时，如果 p[i] 是 p[i−1] 在字母表中的下一个字母，则将 k 加一，否则将 k 置为 1，
	表示重新开始计算连续递增的子串长度。
	然后，用 k 更新 dp[p[i]]的最大值。
	*/ 

	for i, ch := range p {
		if i > 0 {
			if (byte(ch)-p[i-1]+26)%26 == 1 { //1或者-25
				k++
			} else {
				k = 1
			}
		} else {
			k = 1
		}
		dp[ch-'a'] = utility.Max(dp[ch-'a'], k)
	}
	ret := 0
	for _, v := range dp {
		ret += v
	}
	return ret
}
