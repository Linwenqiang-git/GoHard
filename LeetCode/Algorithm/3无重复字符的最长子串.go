package algorithm

import (
	"fmt"
	"strings"
	utility "github.linwenqiang.com/LeetCode/Utility"
)

func lengthOfLongestSubstring(s string) int {
	max := 0
	i := 0
	str_lenght := len(s)
	for i < str_lenght && (str_lenght-i) > max {
		j := i + 1
		child_str := string(s[i])
		for j < str_lenght {
			if !strings.Contains(child_str, string(s[j])) {
				child_str += string(s[j])
				j += 1
			} else {
				if j-i > max {
					max = j - i
				}
				i += 1
				break
			}
		}
		if j == str_lenght {
			if j-i > max {
				max = j - i
			}
			i += 1
		}
	}
	return max
}

func Call_3() {
	fmt.Println(lengthOfLongestSubstring2(" "))
}

//官方解法
func lengthOfLongestSubstring2(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = utility.Max(ans, rk-i+1)
	}
	return ans
}
