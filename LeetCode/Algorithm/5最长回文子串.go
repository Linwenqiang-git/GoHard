package algorithm

import "fmt"

//中心扩展算法
//中心思想：即遍历每个字符，以该字符为扩展中心，向外扩散，直到i-1和j+1不相等时，则之后的都不可能是回文串
func longestPalindrome(s string) string {
	if s == "" {
		return s
	}
	n := len(s)
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, n, i, i)
		//兼容两个相邻字符相同的情况
		left2, right2 := expandAroundCenter(s, n, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
		if start == 0 && end+1 == n {
			break
		}
	}
	return s[start : end+1]
}
func expandAroundCenter(s string, length, Lpoint, Rpoint int) (int, int) {
	for Lpoint >= 0 && Rpoint < length && s[Lpoint] == s[Rpoint] {
		//满足条件继续移动指针
		Lpoint -= 1
		Rpoint += 1
	}
	//表示上一次移动的指针已经不符合条件，移动回上次符合条件的地方
	Lpoint += 1
	Rpoint -= 1
	return Lpoint, Rpoint
}

func Call_5() {
	fmt.Println(longestPalindrome("cbbd"))
}
