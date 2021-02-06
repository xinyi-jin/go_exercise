package main

import "fmt"

/* 几张卡牌 排成一行，每张卡牌都有一个对应的点数。点数由整数数组 cardPoints 给出。

每次行动，你可以从行的开头或者末尾拿一张卡牌，最终你必须正好拿 k 张卡牌。

你的点数就是你拿到手中的所有卡牌的点数之和。

给你一个整数数组 cardPoints 和整数 k，请你返回可以获得的最大点数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-points-you-can-obtain-from-cards
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

// 思路：无论怎么抽取，都是两端的数字 1234 k=3时, 要么是123 要么是12 4 要么是1 34 要么是234 （暴力解会超值）
/* func maxScore(cardPoints []int, k int) int {
	var res int
	for i := 0; i <= k; i++ {
		temp := 0
		for m := 0; m < i; m++ {
			temp += cardPoints[m]
		}
		for n := len(cardPoints) - (k - i); n < len(cardPoints); n++ {
			temp += cardPoints[n]
		}
		if temp > res {
			res = temp
		}
	}
	return res
} */

// 思路二 ：滑动窗口，去除窗口最小值，就是两边抽取数字最大值
func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	// 选取前n-k项作为滑动窗口的最小值
	windowsSize := n - k

	sum := 0
	for _, v := range cardPoints[:windowsSize] {
		sum += v
	}
	minSum := sum
	for i := windowsSize; i < n; i++ {
		// 移动后去除左侧数值，加上右侧数值
		sum += cardPoints[i] - cardPoints[i-windowsSize]
		minSum = min(minSum, sum)
	}
	total := 0
	for _, v := range cardPoints {
		total += v
	}
	return total - minSum
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	cardPoints := make([]int, 7)
	cardPoints = []int{96, 90, 41, 82, 39, 74, 64, 50, 30}
	k := 8
	res := maxScore(cardPoints, k)
	fmt.Println("result:", res)
}
