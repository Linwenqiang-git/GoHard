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
		Lpoint, Rpoint := i, i
		for Lpoint >= 0 && Rpoint < n {
			if s[Lpoint] == s[Rpoint] {
				Lpoint -= 1
				Rpoint += 1
			} else {
				break
			}
		}
		Lpoint += 1
		Rpoint -= 1
		if Rpoint-Lpoint > end-start {
			start, end = Lpoint, Rpoint
		}
		if Lpoint == 0 && Rpoint+1 == n {
			break
		}
	}

	return s[start : end+1]
}

func Call_5() {
	fmt.Println(longestPalindrome("babad"))
}
