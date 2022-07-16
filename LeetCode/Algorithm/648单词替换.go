package algorithm

import (
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	words := strings.Split(sentence, " ")
	rootLen := []int{}
	for _, root := range dictionary {
		rootLen = append(rootLen, len(root))
	}
	for i, v := range words {
		startWithLen := 1<<31 - 1
		for j, length := range rootLen {
			if length <= len(v) {
				startWith := v[0:length]
				if startWith == dictionary[j] && length < startWithLen {
					words[i] = startWith
					startWithLen = length
				}
			}

		}
	}
	return strings.Join(words, " ")
}
