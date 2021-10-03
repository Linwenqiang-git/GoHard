package main

import "fmt"

func isValid(s string) (result bool) {
	if len(s)%2 != 0 {
		return false
	}
	result = true
	stringlen := len(s)
	for i := 0; i < stringlen; i++ {
		if string(s[i]) == ")" || string(s[i]) == "}" || string(s[i]) == "]" {
			return false
		}
		if i+1 < len(s) {
			if string(s[i]) == "(" {
				if string(s[i+1]) == ")" {
					i++ //左右匹配
					continue
				} else {
					//末尾匹配
					if string(s[len(s)-i-1]) == ")" {
						stringlen--
						continue
					} else {
						result = false
						break
					}
				}
			} else if string(s[i]) == "{" {
				if string(s[i+1]) == "}" {
					i++ //左右匹配
					continue
				} else {
					//末尾匹配
					if string(s[len(s)-i-1]) == "}" {
						stringlen--
						continue
					} else {
						result = false
						break
					}
				}
			} else if string(s[i]) == "[" {
				if string(s[i+1]) == "]" {
					i++ //左右匹配
					continue
				} else {
					//末尾匹配
					if string(s[len(s)-i-1]) == "]" {
						stringlen--
						continue
					} else {
						result = false
						break
					}
				}
			}
		}
	}
	return result
}

func main() {
	data := []string{"(([]){})", "({[]})", "}{"}
	for _, value := range data {
		result := isValid(value)
		fmt.Printf("字符串: %s  结果：%v\n", value, result)
	}
}
