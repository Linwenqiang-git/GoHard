package algorithm

func minElements(nums []int, limit int, goal int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	Dvalue := abs(total - goal)
	return (int)((Dvalue + limit - 1) / limit)
}
