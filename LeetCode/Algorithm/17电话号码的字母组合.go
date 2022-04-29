package algorithm

import "fmt"

var digMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

var result = make([]string, 0)

func letterCombinations(digits string) []string {
	if digits == "" {
		return result
	}
	backdata(digits, 0, "")
	return result
}

//这个index对应原字符串的index
func backdata(digits string, index int, splitStr string) {
	if index == len(digits) {
		result = append(result, splitStr)
		return
	}
	currentIndexLen := len(digMap[digits[index]])
	for i := 0; i < currentIndexLen; i++ {
		splitStr += digMap[digits[index]][i]
		backdata(digits, index+1, splitStr)
		splitStr = splitStr[:len(splitStr)-1] //将上一位吐出来
	}
}

func Call_17() {
	fmt.Print(letterCombinations("2349"))
}
