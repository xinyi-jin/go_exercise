package main

import "fmt"

func main() {
	// yuesefu()
}

func yuesefu() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m := 5 //死亡数字
	n := len(arr)
	index := 0   //数字标识
	deadNum := 0 //死亡人数

	for i := 0; i < n; i++ {
		if arr[i] == -1 {
			continue
		}
		arr[i] = index + 1

		if arr[i] == m {
			arr[i] = -1
			index = 0 //重置下标索引
			deadNum++

			if deadNum == n {
				fmt.Println("最后的士兵,index:", i)
				break
			}
		}
		if i == n-1 {
			i = 0
		}
	}
}
