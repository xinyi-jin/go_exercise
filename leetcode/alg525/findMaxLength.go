package alg525

/* 525. 连续数组
给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。

示例 1:
输入: nums = [0,1]
输出: 2
说明: [0, 1] 是具有相同数量0和1的最长连续子数组。

示例 2:
输入: nums = [0,1,0]
输出: 2
说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。

提示：
1 <= nums.length <= 105
nums[i] 不是 0 就是 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/non-decreasing-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

// 官方：前缀和+hash表
func findMaxLength(nums []int) int {
	tmap := make(map[int]int)
	tmap[0] = -1
	sum := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			sum++
		} else {
			sum--
		}
		if pos, ok := tmap[sum]; ok {
			if i-pos > max {
				max = i - pos
			}
		} else {
			tmap[sum] = i
		}
	}
	return max
}
