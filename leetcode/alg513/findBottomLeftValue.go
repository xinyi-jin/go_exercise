package alg513

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*  513. 找树左下角的值
给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。

假设二叉树中至少有一个节点。



示例 1:



输入: root = [2,1,3]
输出: 1
示例 2:



输入: [1,2,3,4,null,5,6,null,null,7]
输出: 7


提示:

二叉树的节点个数的范围是 [1,104]
-231 <= Node.val <= 231 - 1  */

// 思路：深度优先搜索（根据深度来更新最左侧的值）
func findBottomLeftValue(root *TreeNode) int {
	curHeight, res := 0, 0
	var dfs func(node *TreeNode, height int)
	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		height++
		dfs(node.Left, height)
		dfs(node.Right, height)
		if height > curHeight {
			curHeight = height
			res = node.Val
		}
	}
	dfs(root, 0)
	return res
}

// 思路：广度优先搜索（逐个去除队列的右节点和左节点）
func findBottomLeftValueEx(root *TreeNode) int {
	res := 0
	quene := []*TreeNode{root}

	for len(quene) > 0 {
		node := quene[0]
		quene = quene[1:]
		if node.Right != nil {
			quene = append(quene, node.Right)
		}
		if node.Left != nil {
			quene = append(quene, node.Left)
		}
		res = node.Val
	}
	return res
}
