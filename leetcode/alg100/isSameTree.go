package alg100

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* 100. 相同的树
给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。



示例 1：


输入：p = [1,2,3], q = [1,2,3]
输出：true
示例 2：


输入：p = [1,2], q = [1,null,2]
输出：false
示例 3：


输入：p = [1,2,1], q = [1,1,2]
输出：false


提示：

两棵树上的节点数目都在范围 [0, 100] 内
-104 <= Node.val <= 104 */
// 思路：深度优先搜索
func isSameTree(p *TreeNode, q *TreeNode) bool {
	var dfs func(node1, node2 *TreeNode) bool
	dfs = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}
		return dfs(node1.Left, node2.Left) && dfs(node1.Right, node2.Right)
	}
	return dfs(p, q)
}

// 思路：广度优先搜索
func isSameTreeEx(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	quene1, quene2 := []*TreeNode{p}, []*TreeNode{q}
	for len(quene1) > 0 && len(quene2) > 0 {
		node1, node2 := quene1[0], quene2[0]
		quene1, quene2 = quene1[1:], quene2[1:]
		if node1.Val != node2.Val {
			return false
		}
		if node1.Left == nil && node2.Left != nil ||
			node1.Left != nil && node2.Left == nil {
			return false
		}
		if node1.Right == nil && node2.Right != nil ||
			node1.Right != nil && node2.Right == nil {
			return false
		}
		if node1.Left != nil {
			quene1 = append(quene1, node1.Left)
		}
		if node1.Right != nil {
			quene1 = append(quene1, node1.Right)
		}
		if node2.Left != nil {
			quene2 = append(quene2, node2.Left)
		}
		if node2.Right != nil {
			quene2 = append(quene2, node2.Right)
		}
	}
	return true
}
