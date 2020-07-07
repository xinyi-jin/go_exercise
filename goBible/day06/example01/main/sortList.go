package main

import (
	"fmt"
)

func main() {
	var arr = [10]int{1, 5, 6, 2, 3, 1, 0, 10, 78, 9}

	fmt.Println("sort before:", arr)
	//maoPaoSort(arr[:])
	//insertSort(arr[:])
	quickSort(arr[:], 0, len(arr)-1)
	//selectSort(arr[:])
	fmt.Println("sort after:", arr)

}

//冒泡排序
func maoPaoSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr
}

//选择排序
func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] <= arr[j] {
				continue
			}
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}

		}
	}
}

//插入排序
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] <= arr[j] {
				break
			}
			arr[j], arr[j-1] = arr[j-1], arr[j]

		}
	}
}

func quickSort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	val := arr[left]
	k := left

	for i := left; i <= right; i++ {
		if arr[i] < val {
			arr[k] = arr[i]
			arr[i] = arr[k+1]
			k++
		}

		arr[k] = val
		quickSort(arr, left, k-1)
		quickSort(arr, k+1, right)
	}
}
