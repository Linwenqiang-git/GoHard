package algorithm

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		fmt.Println("空字符串问题返回")
		return 0
	}
	result := -1
	needleLen := len(needle)
	haystackLen := len(haystack)
	for index := 0; index < haystackLen; index++ {
		if haystack[index] == needle[0] {
			fmt.Printf("index = %d\n", index)
			if haystackLen-index-1 < needleLen {
				fmt.Printf("长度问题返回.index = %d", index)
				return result
			}
			str := haystack[index : index+needleLen]
			fmt.Printf("截取的字符串:%s\n", str)
			if needle == str {
				fmt.Printf("匹配成功")
				result = index
				break
			}
		} else {
			fmt.Printf("index = %d,%c \n", index, haystack[index])
		}
	}
	return result
}

func Call_28() {
	data := strStr("hello", "ll")
	fmt.Printf("匹配到的元素:%d", data)
}
