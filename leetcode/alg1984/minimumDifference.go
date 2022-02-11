package alg1984

import (
	"sort"
)

/* 1984. 学生分数的最小差值
给你一个 下标从 0 开始 的整数数组 nums ，其中 nums[i] 表示第 i 名学生的分数。另给你一个整数 k 。

从数组中选出任意 k 名学生的分数，使这 k 个分数间 最高分 和 最低分 的 差值 达到 最小化 。

返回可能的 最小差值 。



示例 1：

输入：nums = [90], k = 1
输出：0
解释：选出 1 名学生的分数，仅有 1 种方法：
- [90] 最高分和最低分之间的差值是 90 - 90 = 0
可能的最小差值是 0
示例 2：

输入：nums = [9,4,1,7], k = 2
输出：2
解释：选出 2 名学生的分数，有 6 种方法：
- [9,4,1,7] 最高分和最低分之间的差值是 9 - 4 = 5
- [9,4,1,7] 最高分和最低分之间的差值是 9 - 1 = 8
- [9,4,1,7] 最高分和最低分之间的差值是 9 - 7 = 2
- [9,4,1,7] 最高分和最低分之间的差值是 4 - 1 = 3
- [9,4,1,7] 最高分和最低分之间的差值是 7 - 4 = 3
- [9,4,1,7] 最高分和最低分之间的差值是 7 - 1 = 6
可能的最小差值是 2


提示：

1 <= k <= nums.length <= 1000
0 <= nums[i] <= 105 */

// 思路：先排序，逐个计算相邻k个的首尾两数差值 比较每次结果更新，最终即是所求
func minimumDifference(nums []int, k int) int {
	var res int = -1
	sort.Ints(nums)
	for i := 0; i <= len(nums)-k; i++ {
		temp := nums[i+k-1] - nums[i]
		if res == -1 || temp < res {
			res = temp
		}
	}
	if res == -1 {
		res = 0
	}
	return res
}

// 上述可以优化为此版本 Go性能双百
func minimumDifferenceEx(nums []int, k int) int {
	sort.Ints(nums)
	var res int = nums[len(nums)-1] - nums[0]
	for i := 0; i <= len(nums)-k; i++ {
		if temp := nums[i+k-1] - nums[i]; temp < res {
			res = temp
		}
	}
	return res
}
