package alg70

/* 70. 爬楼梯
简单
相关标签
premium lock icon
相关企业
提示
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？



示例 1：

输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
示例 2：

输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶


提示：

1 <= n <= 45 */

// 思路：动态规划
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	arr := make([]int, n+1)
	arr[1] = 1
	arr[2] = 2
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n]
}

// 动态规划优化内存
func climbStairsEx1(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		tmp := a
		a = b
		b = tmp + a
	}

	return b
}

// 递归 直接超时
func climbStairsEx2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairsEx2(n-1) + climbStairsEx2(n-2)
}

// 记忆化递归
func climbStairsEx3(n int) int {
	nMap := make(map[int]int)
	var df = func(n int) int {
		if n == 1 {
			return 1
		}
		if n == 2 {
			return 2
		}
		if v, ok := nMap[n]; ok {
			return v
		} else {
			nMap[n] = climbStairsEx3(n-1) + climbStairsEx3(n-2)
			return nMap[n]
		}
	}
	return df(n)
}
