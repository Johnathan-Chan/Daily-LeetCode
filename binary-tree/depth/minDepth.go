package depth

// leetcode 111. 二叉树的最小深度 https://leetcode.cn/problems/minimum-depth-of-binary-tree/

import (
	"math"

	"github.com/peaky-lab/Daily-LeetCode/binary-tree/build"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minDepth(root *build.TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	res := math.MaxInt32
	if root.Left != nil {
		res = min(res, minDepth(root.Left))
	}

	if root.Right != nil {
		res = min(res, minDepth(root.Right))
	}

	return res + 1
}
