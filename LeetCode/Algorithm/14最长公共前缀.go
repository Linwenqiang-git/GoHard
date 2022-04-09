package algorithm

import (
	"fmt"
)

/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：2.4 MB, 在所有 Go 提交中击败了9.58% 的用户
*/

func longestCommonPrefix(strs []string) string {
	result := ""
	data := make(map[int]string)
	minLenth := len(strs[0])
	for index, value := range strs {
		data[index] = value
		if minLenth > len(value) {
			minLenth = len(value)
		}
	}
	isExist := false
	for i := 0; i < minLenth; i++ {
		if !isExist {
			value := ""
			for _, v := range data {
				c := v[i]
				if value == "" {
					value = string(v[i])
					continue
				}
				if string(c) != value {
					isExist = true
					break
				}
			}
			if !isExist {

				result += value
			}
		} else {
			break
		}
	}
	return result
}

func Call_14() {
	//strs := []string{"flower", "flow", "flight"}
	strs := []string{"dog", "racecar", "car"}
	result := longestCommonPrefix(strs)
	fmt.Printf("输出点啥吧大哥：%s", result)
}
