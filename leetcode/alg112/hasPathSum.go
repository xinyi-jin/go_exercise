package main

import (
	"fmt"
)

/* 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

func main() {
	arr := []int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 1}
	t := new(TreeNode)

	// 本意是想直接赋值二叉树数组数据，不确定想法具体实现，暂未实现成功
	t.init(arr)

	fmt.Println("TreeNode", t)

	res := hasPathSum(t, 22)
	fmt.Println("res:", res)
}

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 初始化节点
func (t *TreeNode) init(arr []int) *TreeNode {
	n := 0
	for {
		t.Left = createNode(arr[n])
		t.Right = createNode(arr[n+1])
		n++
		if n+1 == len(arr) {
			return t
		}
	}
	return t
}

func createNode(v int) *TreeNode {
	return &TreeNode{Val: v}
}

func (t *TreeNode) setVal(v int) {

	if t == nil {
		fmt.Println("this node not exit")
		return
	}
	t.Val = v //设置节点的Val
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {

	if nil == root {
		return false
	}
	if nil == root.Left && nil == root.Right {
		return sum-root.Val == 0
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
