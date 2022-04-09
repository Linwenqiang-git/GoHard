package algorithm

import (
	"fmt"
)

func isValid(s string) (result bool) {
	stringlen := len(s)
	if stringlen%2 != 0 {
		return false
	}
	matchsymbol := map[byte]byte{
		']': '[',
		')': '(',
		'}': '{',
	}

	stack := []byte{}
	for i := 0; i < stringlen; i++ {
		if matchsymbol[s[i]] > 0 {
			//右侧元素
			if len(stack) == 0 || stack[len(stack)-1] != matchsymbol[s[i]] {
				return false
			}
			//出栈  问题的关键是要把匹配符的括号出栈
			stack = stack[:len((stack))-1]

		} else {
			//左侧元素入栈
			stack = append(stack, s[i])
		}
	}
	//全部匹配成功则为0
	return len(stack) == 0
}

func Call_20() {
	data := []string{"(([]){})", "({[]})", "}{", "([)]"}
	for _, value := range data {
		result := isValid(value)
		fmt.Printf("字符串: %s  结果：%v\n", value, result)
	}
}
