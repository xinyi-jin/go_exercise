package main

import "fmt"

/* 给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。

示例:

输入: [5,2,6,1]
输出: [2,1,1,0]
解释:
5 的右侧有 2 个更小的元素 (2 和 1).
2 的右侧仅有 1 个更小的元素 (1).
6 的右侧有 1 个更小的元素 (1).
1 的右侧有 0 个更小的元素.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */
func main() {
	arr := []int{5, 2, 6, 1}
	res := countSmaller(arr)

	fmt.Println("result:", res)

}

func countSmaller(nums []int) []int {
	l := len(nums)
	resArr := make([]int, l)

	for i := 0; i < l; i++ {
		var conter int
		for j := i; j < l; j++ {
			if i == j {
				continue
			}
			if nums[j] < nums[i] {
				conter++
			}
		}
		resArr[i] = conter
	}
	return resArr
}
