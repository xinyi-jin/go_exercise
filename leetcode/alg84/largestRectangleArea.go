package main

import "fmt"

/* 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。 */

// 84. 柱状图中最大的矩形
// 思路：遍历所有结果，找出其中的最大值 时间复杂度O(n方)，提交后会超时
func largestRectangleAreaEx(heights []int) int {
	l := len(heights)
	maxArea := 0
	if l == 0 {
		return maxArea
	}
	if l == 1 {
		return heights[0]
	}
	for i := 0; i < l; i++ {
		for j := 1; j < l; j++ {
			arr := []int{}
			for x := i; x <= j; x++ {
				if heights[x] == 0 {
					break
				}
				arr = append(arr, heights[x])
			}
			if len(arr) == 0 {
				continue
			}

			curArea := min(arr) * len(arr) * 1
			if curArea > maxArea {
				maxArea = curArea
			}
		}
	}
	return maxArea
}

func min(arr []int) int {
	min := 0
	if len(arr) > 0 {
		min = arr[0]
	}
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

// 思路：升序栈，找当前的位置左边，右边比他小的  下标进行压栈操作，确定可形成矩形面积
// 原始数据，头部和尾部，添加-1，构成整体的升序
func largestRectangleArea(heights []int) int {
	heights = append([]int{-2}, heights...)
	heights = append(heights, -1)

	size := len(heights)
	stack := make([]int, 1, size)

	res := 0
	i := 1
	for i < size {
		// 递增，压栈
		if heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
			continue
		}

		l := len(stack)
		// 当前i标识的是矩形的右边界，stack顶标识的是矩形的左边界
		res = max(res, heights[stack[l-1]]*(i-stack[l-2]-1))

		// 弹栈
		stack = stack[:l-1]
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	arr := []int{1, 1}
	res := largestRectangleArea(arr)
	fmt.Println("res", res)
}
