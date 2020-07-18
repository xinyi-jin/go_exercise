package main

import (
	"fmt"
	"math"
)

/* 一些恶魔抓住了公主（P）并将她关在了地下城的右下角。地下城是由 M x N 个房间组成的二维网格。我们英勇的骑士（K）最初被安置在左上角的房间里，他必须穿过地下城并通过对抗恶魔来拯救公主。

骑士的初始健康点数为一个正整数。如果他的健康点数在某一时刻降至 0 或以下，他会立即死亡。

有些房间由恶魔守卫，因此骑士在进入这些房间时会失去健康点数（若房间里的值为负整数，则表示骑士将损失健康点数）；其他房间要么是空的（房间里的值为 0），要么包含增加骑士健康点数的魔法球（若房间里的值为正整数，则表示骑士将增加健康点数）。

为了尽快到达公主，骑士决定每次只向右或向下移动一步。



编写一个函数来计算确保骑士能够拯救到公主所需的最低初始健康点数。

例如，考虑到如下布局的地下城，如果骑士遵循最佳路径 右 -> 右 -> 下 -> 下，则骑士的初始健康点数至少为 7。

-2 (K)	-3	3
-5	-10	1
10	30	-5 (P)


说明:

骑士的健康点数没有上限。

任何房间都可能对骑士的健康点数造成威胁，也可能增加骑士的健康点数，包括骑士进入的左上角房间以及公主被监禁的右下角房间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/dungeon-game
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

/*
思路来自一个老哥的JavaScript代码

在骑士的房间出发，正向思维很难获得最小血量，但是最后到达公主房间时，骑士血量一定为1。
推导动态规划状态方程，骑士能够走的是右边和下边的屋子，取决于其中最小扣血量的房间。所以从逆向思维来看，骑士逆向进入房间PK后的血量就是正向思维骑士进入房间前PK的血量。二位数组的血量方程应该表达为 dp[i][j] = Math.max(1, Math.min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j])，二维数组也可以进一步优化为一个一维数组不断原来的值，dp[j] = Math.max(1, Math.min(dp[j], dp[j + 1]) - dungeon[i][j]);

作者：an_ger
链接：https://leetcode-cn.com/problems/dungeon-game/solution/yi-wei-shu-zu-ni-xiang-si-wei-dong-tai-gui-hua-nei/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。 */
func main() {
	arr := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}

	result := calculateMinimumHP(arr)
	fmt.Println("result:", result)

}

func calculateMinimumHP(dungeon [][]int) int {
	l, w := len(dungeon), len(dungeon[0])

	// 必须初始化一维数组中的所有元素，指定一个较大的值，然后初始位置w-1 使用1表示最后的血量
	// 初始化dp
	dp := make([]int, w+1)
	for i := 0; i < w+1; i++ {
		dp[i] = math.MaxInt32
	}
	dp[w-1] = 1

	for i := l - 1; i >= 0; i-- {
		for j := w - 1; j >= 0; j-- {
			dp[j] = max(1, min(dp[j], dp[j+1])-dungeon[i][j])
		}
	}

	return dp[0]
}

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
