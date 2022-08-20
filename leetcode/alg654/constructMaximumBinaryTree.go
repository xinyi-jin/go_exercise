package alg654

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var dfs func(nums []int, l, r int) *TreeNode
	dfs = func(nums []int, l, r int) *TreeNode {
		if l > r {
			return nil
		}
		idx := findMaxIdx(nums, l, r)
		node := &TreeNode{}
		node.Val = nums[idx]
		node.Left = dfs(nums, l, idx-1)
		node.Right = dfs(nums, idx+1, r)
		return node
	}

	return dfs(nums, 0, len(nums)-1)
}

func findMaxIdx(nums []int, l, r int) int {
	idx, max := l, nums[l]
	for i := l; i <= r; i++ {
		if nums[i] > max {
			max = nums[i]
			idx = i
		}
	}
	return idx
}
