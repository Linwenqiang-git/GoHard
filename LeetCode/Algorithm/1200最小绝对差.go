package algorithm

import (
	"fmt"
	"sort"
)

func minimumAbsDifference(arr []int) (ret [][]int) {
	sort.Ints(arr)
	ret = make([][]int, len(arr)-1)
	j := 0
	Dvalue := 1<<31 - 1
	for i := 0; i < len(arr); i++ {
		if i+1 < len(arr) && (arr[i+1]-arr[i]) < Dvalue {
			Dvalue = arr[i+1] - arr[i]
			j = 0
		}
		if i+1 < len(arr) && (arr[i+1]-arr[i]) == Dvalue {
			if ret[j] == nil {
				ret[j] = []int{}
			}
			if len(ret[j]) == 0 {
				ret[j] = append(ret[j], arr[i])
				ret[j] = append(ret[j], arr[i+1])
			} else {
				ret[j][0] = arr[i]
				ret[j][1] = arr[i+1]
			}

			j++
		}
	}
	fmt.Print(ret)
	return ret[:j]
}
