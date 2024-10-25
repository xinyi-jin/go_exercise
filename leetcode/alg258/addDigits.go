package alg258

/* 258. 各位相加
给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。返回这个结果。



示例 1:

输入: num = 38
输出: 2
解释: 各位相加的过程为：
38 --> 3 + 8 --> 11
11 --> 1 + 1 --> 2
由于 2 是一位数，所以返回 2。
示例 1:

输入: num = 0
输出: 0


提示：

0 <= num <= 231 - 1 */

// 思路：循环
func addDigits(num int) int {
	if num < 10 {
		return num
	}
	temp := 0
	for ; num > 0; num /= 10 {
		temp += num % 10
	}
	return addDigits(temp)
}

// 递归
func addDigitsEx(num int) int {
	if num < 10 {
		return num
	}
	n, sum := num/10, num%10
	return addDigits(n + sum)
}

// 数学 时间和空间复杂度都为O(1)
func addDigitsExtion(num int) int {
	return (num-1)%9 + 1
}
