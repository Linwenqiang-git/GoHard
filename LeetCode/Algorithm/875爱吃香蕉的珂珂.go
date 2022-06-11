package algorithm

import (
	"math"
	"sort"
)

func minEatingSpeed(piles []int, h int) int {
	heapNum := len(piles)
	sort.Ints(piles)
	if heapNum == h {
		return piles[heapNum-1]
	}
	//每小时最少吃的根数
	mink := avg(piles, h)
	//最大分成的堆数
	allowedSplitNums := heapNum
	if h > allowedSplitNums {
		allowedSplitNums = h
	}
	for allowedSplitNums > 1 {
		splitIndex := 0
		for i, value := range piles {
			if value >= mink {
				splitIndex = i
				break
			}
			allowedSplitNums--
		}
		piles = piles[splitIndex:]
		mink2 := avg(piles, allowedSplitNums)
		if mink2 > mink {
			mink = mink2
		} else {
			return mink
		}
	}
	return mink
}
func avg(p []int, h int) int {
	total := 0.0
	for _, value := range p {
		total += float64(value)
	}
	//每小时最少吃的根数
	mink := int(math.Ceil((total / float64(h))))
	return mink
}

func Call_875() {
	minEatingSpeed([]int{3, 6, 7, 11}, 8)
}
