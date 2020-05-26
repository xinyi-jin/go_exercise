package main

import "fmt"

func main() {
	var arr = [...]int{1, 2, 5, 2, 1, 98, 68, 48, 54, 0}

	fmt.Println("sort before :", arr)

	// insertSort(arr[:])
	// maoPaoSort(arr[:])
	// selectSort(arr[:])
	quickSort(arr[:])

	fmt.Println("sort after :", arr)
}

// 插入排序
func insertSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

// 冒泡排序
func maoPaoSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

// 选择排序
func selectSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				break
			}
		}
	}
}

// 快排
func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	mid := arr[0]
	head, end := 0, len(arr)-1
	for i := 1; i <= end; {
		if arr[i] > mid {
			arr[i], arr[end] = arr[end], arr[i]
			end--
		} else {
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		}
	}
	quickSort(arr[:head])
	quickSort(arr[head+1:])
}
