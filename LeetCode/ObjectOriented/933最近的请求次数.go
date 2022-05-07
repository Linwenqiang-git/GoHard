package objectoriented

type RecentCounter struct {
	pingList []int
}

func Constructor() RecentCounter {
	obj := RecentCounter{
		pingList: make([]int, 0),
	}
	return obj
}

//每次ping的t是严格递增的
func (this *RecentCounter) Ping(t int) int {
	this.pingList = append(this.pingList, t)
	end := t - 3000
	n := len(this.pingList) - 1
	ret := 0
	for n >= 0 {
		if this.pingList[n] >= end {
			ret++
			n--
		} else {
			break
		}
	}
	return ret
}
func (this *RecentCounter) Ping2(t int) int {
	this.pingList = append(this.pingList, t)
	end := t - 3000
	//每次把时间不满足的 直接剔除队列，减小运行时间
	for this.pingList[0] < end {
		this.pingList = this.pingList[1:]
	}
	return len(this.pingList)
}
