package alg69

import (
	"math"
)

/* 69. x 的平方根
简单
相关标签
相关企业
提示
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。



示例 1：

输入：x = 4
输出：2
示例 2：

输入：x = 8
输出：2
解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。


提示：

0 <= x <= 231 - 1 */

// 思路：二分查找
func mySqrt(x int) int {
	pre, suff := 0, x
	res := -1
	for pre <= suff {
		mid := pre + (suff-pre)/2
		n := mid * mid
		if n <= x {
			res = mid
			pre = mid + 1
		} else {
			suff = mid - 1
		}
	}
	return res
}

func mySqrt2(x int) int {
	return int(math.Sqrt(float64(x)))
}

// 逆向思维计算？ 评论区优雅代码
// 变相二分写法
func mySqrt3(x int) int {
	n := x
	for n*n > x {
		n = (n + x/n) / 2
	}
	return n
}
