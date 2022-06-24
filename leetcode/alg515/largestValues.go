package alg515

import "math"

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* 515. 在每个树行中找最大值
给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。



示例1：



输入: root = [1,3,2,5,3,null,9]
输出: [1,3,9]
示例2：

输入: root = [1,2,3]
输出: [1,3]


提示：

二叉树的节点个数的范围是 [0,104]
-231 <= Node.val <= 231 - 1 */

// 思路：广度优先搜索(层序遍历)
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	queen := []*TreeNode{root}
	for len(queen) > 0 {
		maxValue := math.MinInt32
		temp := queen
		queen = nil
		for _, node := range temp {
			maxValue = max(node.Val, maxValue)
			if node.Left != nil {
				queen = append(queen, node.Left)
			}
			if node.Right != nil {
				queen = append(queen, node.Right)
			}
		}
		res = append(res, maxValue)
	}
	return res
}

// 思路：深度优先搜索
func largestValuesEx(root *TreeNode) (ans []int) {
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, curHeight int) {
		if node == nil {
			return
		}
		if curHeight == len(ans) {
			ans = append(ans, node.Val)
		} else {
			ans[curHeight] = max(ans[curHeight], node.Val)
		}
		dfs(node.Left, curHeight+1)
		dfs(node.Right, curHeight+1)
	}
	dfs(root, 0)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
