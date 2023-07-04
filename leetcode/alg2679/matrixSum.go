package alg2679

import "sort"

/* 2679. 矩阵中的和
给你一个下标从 0 开始的二维整数数组 nums 。一开始你的分数为 0 。你需要执行以下操作直到矩阵变为空：

矩阵中每一行选取最大的一个数，并删除它。如果一行中有多个最大的数，选择任意一个并删除。
在步骤 1 删除的所有数字中找到最大的一个数字，将它添加到你的 分数 中。
请你返回最后的 分数 。



示例 1：

输入：nums = [[7,2,1],[6,4,2],[6,5,3],[3,2,1]]
输出：15
解释：第一步操作中，我们删除 7 ，6 ，6 和 3 ，将分数增加 7 。下一步操作中，删除 2 ，4 ，5 和 2 ，将分数增加 5 。最后删除 1 ，2 ，3 和 1 ，将分数增加 3 。所以总得分为 7 + 5 + 3 = 15 。
示例 2：

输入：nums = [[1]]
输出：1
解释：我们删除 1 并将分数增加 1 ，所以返回 1 。


提示：

1 <= nums.length <= 300
1 <= nums[i].length <= 500
0 <= nums[i][j] <= 103 */

// 思路：先排序 后逐个记录每列最大值 就是每次删除的那些数值中的最大值
// 最后返回每列之和即为所求
func matrixSum(nums [][]int) int {
	row, col := len(nums), len(nums[0])
	for i := 0; i < row; i++ {
		sort.Ints(nums[i])
	}
	maxNum := make([]int, col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			maxNum[j] = max(maxNum[j], nums[i][j])
		}
	}
	res := 0
	for _, v := range maxNum {
		res += v
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
