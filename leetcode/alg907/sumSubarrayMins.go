package alg907

import "math"

/* 907. 子数组的最小值之和
中等
698
相关企业
给定一个整数数组 arr，找到 min(b) 的总和，其中 b 的范围为 arr 的每个（连续）子数组。

由于答案可能很大，因此 返回答案模 10^9 + 7 。



示例 1：

输入：arr = [3,1,2,4]
输出：17
解释：
子数组为 [3]，[1]，[2]，[4]，[3,1]，[1,2]，[2,4]，[3,1,2]，[1,2,4]，[3,1,2,4]。
最小值为 3，1，2，4，1，1，2，1，1，1，和为 17。
示例 2：

输入：arr = [11,81,94,43,3]
输出：444


提示：

1 <= arr.length <= 3 * 104
1 <= arr[i] <= 3 * 104 */

// 思路：动态规划（超时）
func sumSubarrayMins(arr []int) int {
	res := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		dp := arr[i]
		for j := i; j < n; j++ {
			dp = int(math.Min(float64(dp), float64(arr[j])))
			res += dp
		}
	}
	return res % (1e9 + 7)
}

// 动态规划 + 单调栈（大到小）
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func sumSubarrayMinsEx(arr []int) (ans int) {
	const mod int = 1e9 + 7
	n := len(arr)
	monoStack := []int{}
	dp := make([]int, n)
	for i, x := range arr {
		for len(monoStack) > 0 && arr[monoStack[len(monoStack)-1]] > x {
			monoStack = monoStack[:len(monoStack)-1]
		}
		k := i + 1
		if len(monoStack) > 0 {
			k = i - monoStack[len(monoStack)-1]
		}
		dp[i] = k * x
		if len(monoStack) > 0 {
			dp[i] += dp[i-k]
		}
		ans = (ans + dp[i]) % mod
		monoStack = append(monoStack, i)
	}
	return
}

/* 作者：力扣官方题解
链接：https://leetcode.cn/problems/sum-of-subarray-minimums/solutions/1929461/zi-shu-zu-de-zui-xiao-zhi-zhi-he-by-leet-bp3k/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。 */
