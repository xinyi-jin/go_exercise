package alg814

/* 814. 二叉树剪枝
给你二叉树的根结点 root ，此外树的每个结点的值要么是 0 ，要么是 1 。

返回移除了所有不包含 1 的子树的原二叉树。

节点 node 的子树为 node 本身加上所有 node 的后代。



示例 1：


输入：root = [1,null,0,0,1]
输出：[1,null,0,null,1]
解释：
只有红色节点满足条件“所有不包含 1 的子树”。 右图为返回的答案。
示例 2：


输入：root = [1,0,1,0,0,0,1]
输出：[1,null,1,null,1]
示例 3：


输入：root = [1,1,0,1,1,0,1,0]
输出：[1,1,0,1,1,null,1]


提示：

树中节点的数目在范围 [1, 200] 内
Node.val 为 0 或 1 */

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：递归
func pruneTree(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)
		if node.Left == nil && node.Right == nil && node.Val == 0 {
			return nil
		}
		return node
	}
	return dfs(root)
}
