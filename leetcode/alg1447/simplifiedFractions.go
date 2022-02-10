package alg1447

import "fmt"

/* 1447. 最简分数
给你一个整数 n ，请你返回所有 0 到 1 之间（不包括 0 和 1）满足分母小于等于  n 的 最简 分数 。分数可以以 任意 顺序返回。



示例 1：

输入：n = 2
输出：["1/2"]
解释："1/2" 是唯一一个分母小于等于 2 的最简分数。
示例 2：

输入：n = 3
输出：["1/2","1/3","2/3"]
示例 3：

输入：n = 4
输出：["1/2","1/3","1/4","2/3","3/4"]
解释："2/4" 不是最简分数，因为它可以化简为 "1/2" 。
示例 4：

输入：n = 1
输出：[]


提示：

1 <= n <= 100 */

// 思路：分子分母分别遍历，校验分子分母是否有相同公因数
func simplifiedFractions(n int) []string {
	res := []string{}
	if n <= 1 {
		return res
	}
	for i := 1; i <= n-1; i++ {
		for j := i + 1; j <= n; j++ {
			if isSimplest(i, j) {
				str := fmt.Sprintf("%d/%d", i, j)
				res = append(res, str)
			}
		}
	}
	return res
}

func isSimplest(a, b int) bool {
	for i := 2; i <= a; i++ {
		if a%i == 0 && b%i == 0 {
			return false
		}
	}
	/* if a%2 == 0 && b%2 == 0 ||
		a%3 == 0 && b%3 == 0 {
		return false
	}
	if a != 1 && b%a == 0 {
		return false
	} */
	return true
}
