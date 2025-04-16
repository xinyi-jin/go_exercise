package alg2537

/* 2537. 统计好子数组的数目
中等
相关标签
相关企业
提示
给你一个整数数组 nums 和一个整数 k ，请你返回 nums 中 好 子数组的数目。

一个子数组 arr 如果有 至少 k 对下标 (i, j) 满足 i < j 且 arr[i] == arr[j] ，那么称它是一个 好 子数组。

子数组 是原数组中一段连续 非空 的元素序列。



示例 1：

输入：nums = [1,1,1,1,1], k = 10
输出：1
解释：唯一的好子数组是这个数组本身。
示例 2：

输入：nums = [3,1,4,3,2,2,4], k = 2
输出：4
解释：总共有 4 个不同的好子数组：
- [3,1,4,3,2,2] 有 2 对。
- [3,1,4,3,2,2,4] 有 3 对。
- [1,4,3,2,2,4] 有 2 对。
- [4,3,2,2,4] 有 2 对。


提示：

1 <= nums.length <= 105
1 <= nums[i], k <= 109 */

// 思路：枚举遍历
// 中级题 会导致超时
func countGood(nums []int, k int) int64 {
	l := len(nums)
	n := int64(0)
	for i := 0; i < l-1; i++ {
		for j := l - 1; j > i; j-- {
			// 得到子数组
			if isGood(nums[i:j+1], k) {
				n++
			}
		}
	}
	return n
}

func isGood(nums []int, k int) bool {
	cntMap := make(map[int]int)
	n := 0
	for _, v := range nums {
		cntMap[v]++
	}
	for _, v := range cntMap {
		if v >= 2 {
			mid := v / 2
			if v%2 != 0 {
				mid = 0
			}
			n += v/2*v - mid
		}
	}
	return n >= k
}

// 官网题解
func countGoodEx(nums []int, k int) int64 {
	n := len(nums)
	same, right := 0, -1
	cnt := make(map[int]int)
	var ans int64 = 0
	for left := 0; left < n; left++ {
		for same < k && right+1 < n {
			right++
			same += cnt[nums[right]]
			cnt[nums[right]]++
		}
		if same >= k {
			ans += int64(n - right)
		}
		cnt[nums[left]]--
		same -= cnt[nums[left]]
	}
	return ans
}
