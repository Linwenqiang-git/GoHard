package algorithm

func findAndReplacePattern(words []string, pattern string) []string {
	ret := []string{}
	for _, word := range words {
		if isMatch(word, pattern) {
			ret = append(ret, word)
		}
	}
	return ret
}

//双射匹配
func isMatch(word string, pattern string) bool {
	mapping := make(map[byte]byte)
	repeat := make(map[byte]bool)
	for i := 0; i < len(word); i++ {
		c, ok := mapping[pattern[i]]
		if !ok {
			mapping[pattern[i]] = word[i]
			_, ok := repeat[word[i]]
			if ok {
				return false
			} else {
				repeat[word[i]] = true
			}
		} else {
			if c != word[i] {
				return false
			}
		}
	}
	return true
}
