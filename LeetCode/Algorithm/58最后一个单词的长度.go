package algorithm

import "fmt"

func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			length++
			continue
		}
		if s[i] == ' ' {
			if length > 0 {
				break
			}
		}
	}
	return length
}

func Call_58() {
	result := lengthOfLastWord("a")
	fmt.Printf("length = %d\n", result)
}
