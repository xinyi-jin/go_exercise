package alg118

/* 118. 杨辉三角
简单
相关标签
premium lock icon
相关企业
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。





示例 1:

输入: numRows = 5
输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
示例 2:

输入: numRows = 1
输出: [[1]]


提示:

1 <= numRows <= 30 */
// 思路：动态规划
func Generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	res := [][]int{{1}, {1, 1}}
	for i := 2; i < numRows; i++ {
		tmp := []int{1}
		for j := 1; j < len(res[i-1]); j++ {
			tmp = append(tmp, res[i-1][j-1]+res[i-1][j])
		}
		tmp = append(tmp, 1)
		res = append(res, tmp)
	}
	return res
}

// 灵神优化 减小内存开销
func GenerateEx(numRows int) [][]int {
	c := make([][]int, numRows)
	for i := range c {
		c[i] = make([]int, i+1)
		c[i][0], c[i][i] = 1, 1
		for j := 1; j < i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
	return c
}
