package main

import "fmt"

/* 给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-the-duplicate-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

// 寻找重复数
func findDuplicate(nums []int) int {
	mapTemp := make(map[int]int, len(nums))

	for _, v := range nums {
		if _, ok := mapTemp[v]; ok {
			return v
		}
		mapTemp[v] = 1
	}
	return -1
}
func main() {
	arr := []int{1, 3, 4, 2, 2}
	fmt.Print(findDuplicate(arr))
}
