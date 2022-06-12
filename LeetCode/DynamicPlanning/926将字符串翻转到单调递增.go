package dynamicprogramming

import (
	utility "github.linwenqiang.com/LeetCode/Utility"
)

func minFlipsMonoIncr(s string) int {
	//dp[i][0] 表示第i个元素是0时，单调递增翻转的最小次数
	//di[i][1] 表示第i个元素是1时，单调递增翻转的最小次数
	/*
		当第i个元素是1时，他前面是0和1都可，所以要取0和1的最小是
		当第i个元素是0时，他前面只能是0
		动态方程：
		II 为示性函数，当事件成立时示性函数值为 1，当事件不成立时示性函数值为 0。
		dp[i][0]​=dp[i−1][0]+I(s[i]=‘1’)
		dp[i][1]​=min(dp[i−1][0],dp[i−1][1])+I(s[i]=‘0’)
	*/
	dp0, dp1 := 0, 0
	for _, c := range s {
		dp0New, dp1New := dp0, utility.Min(dp0, dp1)
		if c == '1' {
			dp0New++
		} else {
			dp1New++
		}
		dp0, dp1 = dp0New, dp1New
	}
	return utility.Min(dp0, dp1)

}
