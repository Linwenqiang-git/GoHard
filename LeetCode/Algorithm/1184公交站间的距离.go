package algorithm

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	AllDistance, ClockwiseDistance := 0, 0
	for value, _ := range distance {
		AllDistance += value
	}
	begin := start
	end := destination
	if begin > end {
		begin = destination
		end = start
	}
	for i := begin; i < end; i++ {
		ClockwiseDistance += distance[i]
	}
	if ClockwiseDistance <= AllDistance-ClockwiseDistance {
		return ClockwiseDistance
	} else {
		return AllDistance - ClockwiseDistance
	}
}
