package alg965

import "fmt"

/* 965. 单值二叉树
如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。

只有给定的树是单值二叉树时，才返回 true；否则返回 false。



示例 1：



输入：[1,1,1,1,1,null,1]
输出：true
示例 2：



输入：[2,2,2,5,2]
输出：false */

//  * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：深度优先搜索
func isUnivalTree(root *TreeNode) bool {
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node != nil {
			if node.Left != nil {
				if node.Left.Val != node.Val {
					return false
				}
				if !dfs(node.Left) {
					return false
				}
			}
			if node.Right != nil {
				if node.Right.Val != node.Val {
					return false
				}
				if !dfs(node.Right) {
					return false
				}
			}
		}
		return true
	}
	return dfs(root)
}

func convertTree(nums []int) *TreeNode {
	var df func(nums []int, i, c int) *TreeNode
	df = func(nums []int, i, c int) *TreeNode {
		if i > len(nums) {
			return nil
		}
		val := nums[i-1]
		if val == -1 && i != len(nums) {
			val = nums[i]
		}
		tree := &TreeNode{Val: val}
		tree.Left = df(nums, 2*i+c, c)
		tree.Right = df(nums, 2*i+1+c, c)
		tree = nil
		return tree
	}
	return df(nums, 1, 0)
}

func traverTree(tree *TreeNode) {
	var df func(node *TreeNode)
	df = func(node *TreeNode) {
		if node == nil {
			return
		}
		df(node.Left)
		df(node.Right)
		fmt.Printf("%v", node.Val)
	}
	df(tree)
}
