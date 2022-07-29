package alg593

/* 593. 有效的正方形
给定2D空间中四个点的坐标 p1, p2, p3 和 p4，如果这四个点构成一个正方形，则返回 true 。

点的坐标 pi 表示为 [xi, yi] 。输入 不是 按任何顺序给出的。

一个 有效的正方形 有四条等边和四个等角(90度角)。



示例 1:

输入: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
输出: True
示例 2:

输入：p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,12]
输出：false
示例 3:

输入：p1 = [1,0], p2 = [-1,0], p3 = [0,1], p4 = [0,-1]
输出：true


提示:

p1.length == p2.length == p3.length == p4.length == 2
-104 <= xi, yi <= 104 */

// 思路：判断是否是等腰三角（勾股定理）
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	return isEquicruralTriangle(p1, p2, p3) && isEquicruralTriangle(p2, p3, p4) &&
		isEquicruralTriangle(p1, p3, p4) && isEquicruralTriangle(p1, p2, p4)
}

func isEquicruralTriangle(p1, p2, p3 []int) bool {
	c1 := (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
	c2 := (p1[0]-p3[0])*(p1[0]-p3[0]) + (p1[1]-p3[1])*(p1[1]-p3[1])
	c3 := (p2[0]-p3[0])*(p2[0]-p3[0]) + (p2[1]-p3[1])*(p2[1]-p3[1])
	return c1 > c2 && c2 == c3 && c1 == c2+c3 ||
		c2 > c3 && c1 == c3 && c2 == c1+c3 ||
		c3 > c1 && c1 == c2 && c3 == c1+c2
}
