package levelorder

import "github.com/peaky-lab/Daily-LeetCode/binary-tree/build"

// leetcode 102. 二叉树的层序遍历 https://leetcode.cn/problems/binary-tree-level-order-traversal/

func levelOrder(root *build.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*build.TreeNode{root}
	for len(queue) > 0 {
		count := len(queue)
		ans := []int{}
		for count > 0 {
			count--
			ans = append(ans, queue[0].Val)
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		res = append(res, ans)
	}
	return res
}
