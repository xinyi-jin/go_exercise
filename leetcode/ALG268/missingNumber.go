package alg268

import "sort"

/* 268. 丢失的数字
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。



示例 1：

输入：nums = [3,0,1]
输出：2
解释：n = 3，因为有 3 个数字，所以所有的数字都在范围 [0,3] 内。2 是丢失的数字，因为它没有出现在 nums 中。
示例 2：

输入：nums = [0,1]
输出：2
解释：n = 2，因为有 2 个数字，所以所有的数字都在范围 [0,2] 内。2 是丢失的数字，因为它没有出现在 nums 中。
示例 3：

输入：nums = [9,6,4,2,3,5,7,0,1]
输出：8
解释：n = 9，因为有 9 个数字，所以所有的数字都在范围 [0,9] 内。8 是丢失的数字，因为它没有出现在 nums 中。
示例 4：

输入：nums = [0]
输出：1
解释：n = 1，因为有 1 个数字，所以所有的数字都在范围 [0,1] 内。1 是丢失的数字，因为它没有出现在 nums 中。


提示：

n == nums.length
1 <= n <= 104
0 <= nums[i] <= n
nums 中的所有数字都 独一无二 */

/*
	解题思路：
	先排序，后查找。
*/
func missingNumber(nums []int) int {
	sort.Sort(SortBySliceInt(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}

/*
	解题思路：
	求和:性能较为稳定
*/
func missingNumberEx(nums []int) int {
	n := len(nums)
	sum := (n + 1) * n / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}

/*
	解题思路：
	位运算
*/
func missingNumberExt(nums []int) int {
	num := len(nums)
	for i := 0; i < len(nums); i++ {
		num ^= nums[i] ^ i
	}
	return num
}

/*
	解题思路：
	最优解：大概意思就是就算对应下标位置的排序后值与实际值的差
*/
func missingNumberExti(nums []int) int {
	miss := 0
	var i int
	for i = 0; i < len(nums); i++ {
		miss += i
		miss -= nums[i]
	}
	return miss + i

}

type SortBySliceInt []int

func (a SortBySliceInt) Len() int           { return len(a) }
func (a SortBySliceInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBySliceInt) Less(i, j int) bool { return a[i] < a[j] }
