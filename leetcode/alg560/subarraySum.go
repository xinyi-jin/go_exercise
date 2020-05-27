package main

import "fmt"

/* 给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

示例 1 :

输入:nums = [1,1,1], k = 2
输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subarray-sum-equals-k
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

func subarraySum(nums []int, k int) int {
	if len(nums) == 1 && nums[0] == k {
		return 1
	}
	sum, count := 0, 0
	for i := 0; i < len(nums); i++ {
		for _, v := range nums {
			sum += v
			if sum == k {
				count++
				break
			}
		}
		if nums[i] == k {
			count++
		}
		// 重置循环条件
		i = 0
		sum = 0
		nums = nums[1:]
	}
	return count
}

func main() {
	str := []int{1}
	res := subarraySum(str, 1)
	fmt.Println(res)
}
