package alg883

/* 883. 三维形体投影面积
在 n x n 的网格 grid 中，我们放置了一些与 x，y，z 三轴对齐的 1 x 1 x 1 立方体。

每个值 v = grid[i][j] 表示 v 个正方体叠放在单元格 (i, j) 上。

现在，我们查看这些立方体在 xy 、yz 和 zx 平面上的投影。

投影 就像影子，将 三维 形体映射到一个 二维 平面上。从顶部、前面和侧面看立方体时，我们会看到“影子”。

返回 所有三个投影的总面积 。



示例 1：



输入：[[1,2],[3,4]]
输出：17
解释：这里有该形体在三个轴对齐平面上的三个投影(“阴影部分”)。
示例 2:

输入：grid = [[2]]
输出：5
示例 3：

输入：[[1,0],[0,2]]
输出：8


提示：

n == grid.length == grid[i].length
1 <= n <= 50
0 <= grid[i][j] <= 50 */

/* 思路
	根据坐标推算每个点正方体数量，如下所示：
	[1,2],    ==>数量代入坐标 00 01  ==>即可得到shadow.png图示三维立体图形数据
	[3,4]			         10 11

然后就是处理三视图数据了 */

// 执行用时: 0 ms
// 内存消耗: 3.5 MB  
func projectionArea(grid [][]int) int {
	res := 0
	n := len(grid)
	for i := 0; i < n; i++ {
		x, y, z := 0, 0, 0
		for _, v := range grid[i] {
			y = max(v, y)
			if v > 0 {
				z++
			}
		}
		for j := 0; j < n; j++ {
			x = max(grid[j][i], x)
		}
		res += (x + y + z)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
