package main

import "fmt"

func main() {
	result := twoSum2([]int{2, 3, 4}, 6)
	fmt.Println("result:", result)

}

/* 给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。

函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。

说明:

返回的下标值（index1 和 index2）不是从零开始的。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
示例:

输入: numbers = [2, 7, 11, 15], target = 9
输出: [1,2]
解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

// 使用快慢指针暴力求解，时间复杂度O(n方)
func twoSum(numbers []int, target int) []int {
	res := make([]int, 2)
	l := len(numbers)
	if l < 2 {
		return res
	}

	// 使用快慢指针实现
	for i := 0; i < l-1; i++ {
		currentNum := target - numbers[i]
		for j := i + 1; j < l; j++ {
			if currentNum == numbers[j] {
				res[0] = i + 1
				res[1] = j + 1
			}
		}
	}
	return res
}

// 双指针法，因为是有序排列的数组，时间复杂度O(N)
func twoSum1(numbers []int, target int) []int {
	low, high := 0, len(numbers)-1
	for low < high {
		sum := numbers[low] + numbers[high]
		if sum == target {
			return []int{low + 1, high + 1}
		} else if sum < target {
			// 当前两数的和若比目标值小，证明需要增大其中一个数字，high表示最大，把low位的数字右移一位就好了
			low++
		} else {
			high--
		}
	}
	return []int{}
}

// 二分查找法，因为是有序排列的数组，时间复杂度O(n*log n)
func twoSum2(numbers []int, target int) []int {
	res := make([]int, 2)
	l := len(numbers)
	if l < 2 {
		return res
	}

	// 使用快慢指针实现
	for i := 0; i < l-1; i++ {
		currentNum := target - numbers[i]
		// 二分法查找
		low, high := i+1, l-1
		for low <= high {
			mid := (high-low)/2 + low
			if currentNum == numbers[mid] {
				res[0] = i + 1
				res[1] = mid + 1
			}
			if currentNum < numbers[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return res
}
