package alg1020

/* 1020. 飞地的数量
给你一个大小为 m x n 的二进制矩阵 grid ，其中 0 表示一个海洋单元格、1 表示一个陆地单元格。

一次 移动 是指从一个陆地单元格走到另一个相邻（上、下、左、右）的陆地单元格或跨过 grid 的边界。

返回网格中 无法 在任意次数的移动中离开网格边界的陆地单元格的数量。



示例 1：


输入：grid = [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
输出：3
解释：有三个 1 被 0 包围。一个 1 没有被包围，因为它在边界上。
示例 2：


输入：grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
输出：0
解释：所有 1 都在边界上或可以到达边界。


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 500
grid[i][j] 的值为 0 或 1 */

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 官方dfs题解
func numEnclaves(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 || vis[r][c] {
			return
		}
		vis[r][c] = true
		for _, d := range dirs {
			dfs(r+d.x, c+d.y)
		}
	}
	for i := range grid {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 1; j < n-1; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 && !vis[i][j] {
				ans++
			}
		}
	}
	return
}
