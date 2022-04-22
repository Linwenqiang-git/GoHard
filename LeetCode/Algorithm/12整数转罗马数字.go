package algorithm
import "fmt"

func intToRoman(num int) string {
	var romanMapping = make(map[int]string,0)
	romanMapping[1] = "I"
	romanMapping[4] = "IV"
	romanMapping[5] = "V"
	romanMapping[9] = "IX"
	romanMapping[10] = "X"
	romanMapping[40] = "XL"
	romanMapping[50] = "L"
	romanMapping[90] = "XC"
	romanMapping[100] = "C"
	romanMapping[400] = "CD"
	romanMapping[500] = "D"
	romanMapping[900] = "CM"
	romanMapping[1000] = "M"
	romanKeys := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	result := ""
	start := 0
	for num > 0 && start < 13{
		bit := num / romanKeys[start]
		if bit > 0{
			value,ok := romanMapping[bit * romanKeys[start]]
			bitNum := bit * romanKeys[start]
			if ok{
				result += value
			}else{
				for bit >0{
					result+=romanMapping[romanKeys[start]]
					bit--
				}
			}
			num -= bitNum
		}
		start++
	}
	return result
}

func Call_12() {
	fmt.Println(intToRoman(3999))
}