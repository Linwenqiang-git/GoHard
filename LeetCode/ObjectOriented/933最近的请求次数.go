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
			ret++type MovingAverage struct {
    size int
    ori_data []int
}


/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
    return MovingAverage{
        size:size,
        data:make([]int,0)
    }
}


func (this *MovingAverage) Next(val int) float64 {
	this.ori_data = append(this.ori_data,val)
	ori_lenth = len(this.ori_data)
	if ori_data >=3{
		sum := this.ori_data[ori_lenth-1] + this.ori_data[ori_lenth - 2] + this.ori_data[ori_lenth - 3]
		return float64(sum) /3
	}else{
		sum := 0
		for _,v := range this.ori_data{
			sum += v
		}
		return float64(sum) / ori_lenth
	}
}
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
