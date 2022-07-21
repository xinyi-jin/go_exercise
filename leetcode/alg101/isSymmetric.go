package alg101

/* 101. 对称二叉树
给你一个二叉树的根节点 root ， 检查它是否轴对称。



示例 1：


输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：


输入：root = [1,2,2,null,3,null,3]
输出：false


提示：

树中节点数目在范围 [1, 1000] 内
-100 <= Node.val <= 100


进阶：你可以运用递归和迭代两种方法解决这个问题吗？ */

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
func isSymmetric(root *TreeNode) bool {
	var dfs func(p, q *TreeNode) bool
	dfs = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		return p.Val == q.Val && dfs(p.Left, q.Right) && dfs(p.Right, q.Left)
	}
	return dfs(root, root)
}

// BFS
func isSymmetricEx(root *TreeNode) bool {
	quene1, quene2 := []*TreeNode{root}, []*TreeNode{root}
	for len(quene1) > 0 && len(quene2) > 0 {
		node1, node2 := quene1[0], quene2[0]
		quene1, quene2 = quene1[1:], quene2[1:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil && node2 != nil ||
			node1 != nil && node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}
		quene1 = append(quene1, node1.Left)
		quene1 = append(quene1, node1.Right)
		quene2 = append(quene2, node2.Right)
		quene2 = append(quene2, node2.Left)
	}
	return true
}
