package interviewquestion

import "math"

func oneEditAway(first string, second string) bool {
	firstLen := len(first)
	secondLen := len(second)
	if math.Abs(float64(firstLen)-float64(secondLen)) > 1 {
		return false
	}
	//挖坑点：对于长度为0的字符串，使用索引会报错
	if firstLen == 0 || secondLen == 0 {
		return true
	}
	operateCount := 0
	if math.Abs(float64(firstLen)-float64(secondLen)) == 1 {
		//添加或者删除一个
		i, j := 0, 0
		for i < firstLen || j < secondLen {
			if (i == firstLen && j < secondLen) || (i < firstLen && j == secondLen) {
				if operateCount == 1 {
					return false
				} else {
					return true
				}
			}
			if first[i] != second[j] {
				if operateCount == 0 {
					//模拟将长的那个删除一个
					if firstLen > secondLen {
						i++
					} else {
						j++
					}
					operateCount++
					continue
				} else {
					return false
				}
			}
			if i < firstLen {
				i++
			}
			if j < secondLen {
				j++
			}
		}
	}
	if firstLen == secondLen {
		//替换操作
		for i := 0; i < firstLen; i++ {
			if first[i] != second[i] {
				if operateCount == 0 {
					//模拟一次替换操作
					operateCount++
				} else {
					return false
				}
			}
		}
	}
	return true
}
