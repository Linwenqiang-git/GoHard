package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)
	var maxMinus int
	initmaxMinus := false
	integer := 0
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			if !initmaxMinus {
				maxMinus = nums[i]
				initmaxMinus = true
				continue
			}
			if nums[i] > maxMinus {
				maxMinus = nums[i]
			}
		} else if nums[i] >= 0 && integer == 0 {
			//找到第一个正数，开始计算
			initMax := nums[i]
			max := initMax
			for j := i + 1; j < n; j++ {
				newNum := nums[j] + initMax
				//当前的数比累加的还要大,那就从头累加
				if newNum < nums[j] {
					initMax = nums[j]
					if initMax >= max {
						max = initMax
					}

					//fmt.Printf("initMax = %d\n", initMax)
					continue
				}
				if newNum >= max {
					max = newNum
				}

				initMax = newNum
				fmt.Printf("nums[j] = %d,initMax = %d\n", nums[j], initMax)
			}
			fmt.Println("从这里返回1")
			return max
		}
	}
	//fmt.Println("从这里返回2")
	return maxMinus
}

//官方解法
func maxSubArray2(nums []int) int {
	//若当前指针所指元素之前的和小于0，则丢弃当前元素之前的数列
	max := nums[0] //最大和
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func main() {
	//nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 1}
	nums := []int{1, -2, 0}
	result := maxSubArray(nums)
	fmt.Printf("最大结果：%d \n", result)
}
