package slot

import (
	"sort"
)

func IsLine(data []int) (e int, count int) {
	tempData := make([]int, 5)
	copy(tempData, data)
	dataMap := make(map[int]int, 5)

	for _, v := range tempData {
		if _, ok := dataMap[v]; !ok {
			dataMap[v] = 1
		} else {
			dataMap[v]++
		}
	}
	for _, v := range tempData {
		if dataMap[v] >= 3 {
			return v, dataMap[v]
		}
	}
	return -1, -1
}

// 处理单个结果
func IsLineEx(data []int) (e int, count int) {
	tempData := make([]int, 5)
	copy(tempData, data)
	sort.Ints(tempData)

	curPos := 0
	curElement := tempData[0]

	for i := 1; i < len(tempData); i++ {
		n := i - curPos
		if curElement == tempData[len(tempData)-1] && len(tempData)-curPos >= 3 {
			return curElement, len(tempData) - curPos
		}
		if n >= 3 {
			if curElement != tempData[i] {
				return curElement, n
			}
		} else {
			if curElement != tempData[i] {
				curElement = tempData[i]
				curPos = i
			}
		}
	}
	return -1, -1
}

// 可处理多个结果
func IsLineExtend(data []int, flag bool) (e int, count int) {
	tempData := make([]int, 5)
	copy(tempData, data)
	sort.Ints(tempData)

	curPos := 0
	curElement := tempData[0]
	arrResult := make([]int64, 0)
	for i := 1; i < len(tempData); i++ {
		n := i - curPos
		if curElement == tempData[len(tempData)-1] && len(tempData)-curPos >= 3 {
			if flag {
				calcValue := int64(curElement) + int64(n)*10000
				arrResult = append(arrResult, calcValue)
				break
			}
			return curElement, len(tempData) - curPos
		}
		if n >= 3 {
			if curElement != tempData[i] {
				if flag {
					calcValue := int64(curElement) + int64(n)*10000
					arrResult = append(arrResult, calcValue)

					curElement = tempData[i]
					curPos = i
					continue
				}
				return curElement, n
			}
		} else {
			if curElement != tempData[i] {
				curElement = tempData[i]
				curPos = i
			}
		}
	}
	if flag && len(arrResult) > 0 {
		// 处理多个结果
	}
	return -1, -1
}
