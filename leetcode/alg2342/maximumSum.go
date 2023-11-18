package alg2342

import (
	"sort"
)

/* 2342. 数位和相等数对的最大和
提示
中等
34
相关企业
给你一个下标从 0 开始的数组 nums ，数组中的元素都是 正 整数。请你选出两个下标 i 和 j（i != j），且 nums[i] 的数位和 与  nums[j] 的数位和相等。

请你找出所有满足条件的下标 i 和 j ，找出并返回 nums[i] + nums[j] 可以得到的 最大值 。



示例 1：

输入：nums = [18,43,36,13,7]
输出：54
解释：满足条件的数对 (i, j) 为：
- (0, 2) ，两个数字的数位和都是 9 ，相加得到 18 + 36 = 54 。
- (1, 4) ，两个数字的数位和都是 7 ，相加得到 43 + 7 = 50 。
所以可以获得的最大和是 54 。
示例 2：

输入：nums = [10,12,19,14]
输出：-1
解释：不存在满足条件的数对，返回 -1 。


提示：

1 <= nums.length <= 105
1 <= nums[i] <= 109 */

type SortByNum []int

func (a SortByNum) Len() int           { return len(a) }
func (a SortByNum) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByNum) Less(i, j int) bool { return a[i] > a[j] }

// 思路：使用数据构建缓存，减小性能消耗
// 使用数据结构算法：数组、哈希、排序
func maximumSum(nums []int) int {
	res := -1
	// 排序后，第一次查找到的即为最大值
	sort.Sort(SortByNum(nums))

	n := len(nums)
	cash := make([]int, n)
	cash[0] = calcSum(nums[0])
	findMap := make(map[int]bool)
	for i := 0; i < n; i++ {
		if cash[i] == 0 {
			cash[i] = calcSum((nums[i]))
		}
		// 跳过已查找到的
		if findMap[cash[i]] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if cash[j] == 0 {
				cash[j] = calcSum((nums[j]))
			}
			if cash[i] == cash[j] {
				// 记录已查找到的
				findMap[cash[i]] = true
				// 更新最大值
				res = max(res, nums[i]+nums[j])
				break
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 计算位数和
func calcSum(num int) int {
	res := 0
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return res
}
