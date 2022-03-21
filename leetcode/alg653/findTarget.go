package alg653

import "fmt"

/* 653. 两数之和 IV - 输入 BST
给定一个二叉搜索树 root 和一个目标结果 k，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。



示例 1：


输入: root = [5,3,6,2,4,null,7], k = 9
输出: true
示例 2：


输入: root = [5,3,6,2,4,null,7], k = 28
输出: false


提示:

二叉树的节点个数的范围是  [1, 104].
-104 <= Node.val <= 104
root 为二叉搜索树
-105 <= k <= 105 */

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	var dfs func(node *TreeNode) bool
	set := make(map[int]int)
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if _, ok := set[k-node.Val]; ok {
			return true
		}
		set[node.Val] = 0
		return dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}

func convertTree(nums []int) *TreeNode {
	var df func(nums []int, i int) *TreeNode
	df = func(nums []int, i int) *TreeNode {
		if i > len(nums) {
			return nil
		}
		tree := &TreeNode{Val: nums[i-1]}
		tree.Left = df(nums, 2*i)
		tree.Right = df(nums, 2*i+1)
		return tree
	}
	return df(nums, 1)
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
