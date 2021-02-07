package main

import "fmt"

/* 给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。

我们是这样定义一个非递减数列的： 对于数组中所有的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/non-decreasing-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

// 思路 一开始只想到了了对比两个数值，没有想到如果l>r的时候，l还大于r后边的数 这种情况
// 对比官方的解答，使用for range 去进行遍历操作时间上提升了8ms，内存消耗和官方一致。
// 但是需要注意不能直接使用for range 里的v进行改变切片值。因为是v是copy出来的切片值，不会改变原始切片值
func checkPossibility(nums []int) bool {
	n := len(nums)
	if n > 2 {
		cnt := 0
		for i, v := range nums[:n-1] {
			l, r := v, nums[i+1]
			if l > r {
				cnt++
				if cnt > 1 {
					return false
				}
				if i > 0 && nums[i-1] > r {
					nums[i+1] = l
				}
			}
		}
	}
	return true
}

// 官方
/* func checkPossibility(nums []int) bool {
	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			cnt++
			if cnt > 1 {
				return false
			}
			if i > 0 && y < nums[i-1] {
				nums[i+1] = x
			}
		}
	}
	return true
} */

func main() {
	nums := make([]int, 10)
	nums = []int{3, 4, 2, 3}
	res := checkPossibility(nums)
	fmt.Println("result:", res)
}
