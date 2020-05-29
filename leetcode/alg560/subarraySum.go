package main

import "fmt"

/* 给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

示例 1 :

输入:nums = [1,1,1], k = 2
输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subarray-sum-equals-k
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

// 560. 和为K的子数组
func subarraySum(nums []int, k int) int {
	count := 0 //计数
	for {
		sum := 0 //统计目前连续数组的和
		for i := 0; i < len(nums); i++ {
			sum += nums[i]
			if sum == k {
				count++
			}
		}
		// 退出循环
		if len(nums) == 1 {
			break
		}
		nums = nums[1:]
	}
	return count
}

func main() {
	// str := []int{1, 1, 1, 1}
	str := []int{28, 54, 7, -70, 22, 65, -6}
	res := subarraySum(str, 100)
	fmt.Println(res)
}
