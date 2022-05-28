package depth

import "github.com/peaky-lab/Daily-LeetCode/binary-tree/build"

// leetcode 104. 二叉树的最大深度 https://leetcode.cn/problems/maximum-depth-of-binary-tree/

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func dfs(root *build.TreeNode, depth int) int {
	if root == nil {
		return 0
	}

	return max(dfs(root.Left, depth)+1, dfs(root.Right, depth)+1)
}

func maxDepth(root *build.TreeNode) int {
	if root == nil {
		return 0
	}

	return dfs(root, 0)
}
