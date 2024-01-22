package alg670

/* 670. 最大交换
中等
相关标签
相关企业
给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

示例 1 :

输入: 2736
输出: 7236
解释: 交换数字2和数字7。
示例 2 :

输入: 9973
输出: 9973
解释: 不需要交换。
注意:

给定数字的范围是 [0, 108] */

// 数组
func maximumSwap(num int) int {
	if num < 10 {
		return num
	}
	arr := make([]int, 0)
	for num > 0 {
		n := num % 10
		num /= 10
		arr = append(arr, n)
	}
	if len(arr) >= 2 {
		exchanged := false
		for i := len(arr) - 1; i >= 0; i-- {
			if arr[i] == 9 {
				continue
			}
			maxNum, maxPos := arr[i], i
			for j := i - 1; j >= 0; j-- {
				if maxNum < arr[j] {
					maxNum = arr[j]
					maxPos = j
					exchanged = true
				}
				if exchanged {
					if maxNum == arr[j] {
						maxNum = arr[j]
						maxPos = j
					}
				}
			}
			if exchanged {
				arr[i], arr[maxPos] = arr[maxPos], arr[i]
				break
			}
		}
	}
	res := arr[0]
	scale := 10
	for i := 1; i < len(arr); i++ {
		res += scale * arr[i]
		scale *= 10
	}
	return res
}
