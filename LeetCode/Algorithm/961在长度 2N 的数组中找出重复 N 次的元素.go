package algorithm

import (
	"sort"
	"strings"
)

func repeatedNTimes(nums []int) int {
	sort.Ints(nums)
	ret := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[i+1] {
			ret = nums[i]
			break
		}
	}
	return ret
}

func numUniqueEmails(emails []string) int {
	email := make(map[string]bool)
	for _, value := range emails {
		i := strings.IndexByte(value, '@')
		local := strings.SplitN(value[:i], "+", 2)[0] // 去掉本地名第一个加号之后的部分
		local = strings.ReplaceAll(local, ".", "")    // 去掉本地名中所有的句点
		email[local+value[i:]] = true
	}

	return len(email)
}
