package alg1380

/* 1380. 矩阵中的幸运数
给你一个 m * n 的矩阵，矩阵中的数字 各不相同 。请你按 任意 顺序返回矩阵中的所有幸运数。

幸运数是指矩阵中满足同时下列两个条件的元素：

在同一行的所有元素中最小
在同一列的所有元素中最大


示例 1：

输入：matrix = [[3,7,8],[9,11,13],[15,16,17]]
输出：[15]
解释：15 是唯一的幸运数，因为它是其所在行中的最小值，也是所在列中的最大值。
示例 2：

输入：matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
输出：[12]
解释：12 是唯一的幸运数，因为它是其所在行中的最小值，也是所在列中的最大值。
示例 3：

输入：matrix = [[7,8],[1,2]]
输出：[7]


提示：

m == mat.length
n == mat[i].length
1 <= n, m <= 50
1 <= matrix[i][j] <= 10^5
矩阵中的所有元素都是不同的 */

// 思路：逐个求出每行的最小值 与 每列的最大值
// 因为元素数值各不相同，随意最后相同数字个数>1的即满足两个条件
func luckyNumbers(matrix [][]int) (res []int) {
	numMap := make(map[int]int)
	for i := 0; i < len(matrix); i++ {
		min := matrix[i][0]
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] < min {
				min = matrix[i][j]
			}
		}
		numMap[min]++
	}
	for i := 0; i < len(matrix[0]); i++ {
		max := matrix[0][i]
		for j := 1; j < len(matrix); j++ {
			if matrix[j][i] > max {
				max = matrix[j][i]
			}
		}
		numMap[max]++
	}
	for k, v := range numMap {
		if v > 1 {
			res = append(res, k)
		}
	}
	return
}
