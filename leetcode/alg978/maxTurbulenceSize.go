package main

import (
	"fmt"
)

/* 当 A 的子数组 A[i], A[i+1], ..., A[j] 满足下列条件时，我们称其为湍流子数组：

若 i <= k < j，当 k 为奇数时， A[k] > A[k+1]，且当 k 为偶数时，A[k] < A[k+1]；
或 若 i <= k < j，当 k 为偶数时，A[k] > A[k+1] ，且当 k 为奇数时， A[k] < A[k+1]。
也就是说，如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是湍流子数组。

返回 A 的最大湍流子数组的长度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-turbulent-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

// 本题采用官方题解
func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	ans := 1
	left, right := 0, 0
	for right < n-1 {
		if left == right {
			if arr[left] == arr[left+1] {
				left++
			}
			right++
		} else {
			if arr[right-1] < arr[right] && arr[right] > arr[right+1] {
				right++
			} else if arr[right-1] > arr[right] && arr[right] < arr[right+1] {
				right++
			} else {
				left = right
			}
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	arr := make([]int, 10)
	arr = []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
	res := maxTurbulenceSize(arr)
	fmt.Println("result:", res)
}
