package alg111

import "math"

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* 111. 二叉树的最小深度
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。



示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：2
示例 2：

输入：root = [2,null,3,null,4,null,5,null,6]
输出：5


提示：

树中节点数的范围在 [0, 105] 内
-1000 <= Node.val <= 1000 */

// 思路：DFS
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}
	return minD + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// BFS
func minDepthEx(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queen := []*TreeNode{root}
	count := []int{1}
	for i := 0; i < len(queen); i++ {
		node := queen[i]
		depth := count[i]
		if node.Left == nil && node.Right == nil {
			return depth
		}
		if node.Left != nil {
			queen = append(queen, node.Left)
			count = append(count, depth+1)
		}
		if node.Right != nil {
			queen = append(queen, node.Right)
			count = append(count, depth+1)
		}
	}
	return 0
}
