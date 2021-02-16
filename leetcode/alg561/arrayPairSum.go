package main

import (
	"fmt"
	"sort"
)

/* 给定长度为 2n 的整数数组 nums ，你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从 1 到 n 的 min(ai, bi) 总和最大。

返回该 最大总和 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/array-partition-i
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

// 思路：升序排列后，每两个相邻的数字小的那个就是相加的最大值，顺序去n个数字就是相加的最小值
func arrayPairSum(nums []int) int {
	var res int
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return res
}

func main() {
	var nums []int
	nums = make([]int, 10)
	nums = []int{1, 2, 3, 4}
	res := arrayPairSum(nums)
	fmt.Println("result :", res)
}
