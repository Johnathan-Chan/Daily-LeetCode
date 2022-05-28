package postorder

import "github.com/peaky-lab/Daily-LeetCode/binary-tree/build"

/*
leetcode 145. 二叉树的后序遍历 https://leetcode.cn/problems/binary-tree-postorder-traversal/
binary tree inorder
recursion and stack iteration
*/

// recursion
func dfs(root *build.TreeNode, res *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left, res)
	dfs(root.Right, res)
	*res = append(*res, root.Val)
}

func postorder_recursion(root *build.TreeNode) []int {

	if root == nil {
		return []int{}
	}

	res := []int{}
	dfs(root, &res)
	return res
}

// stack iteration
func postorder_stack_iteration(root *build.TreeNode) []int {

	if root == nil {
		return []int{}
	}

	res := []int{}
	stack := []*build.TreeNode{root}

	for len(stack) > 0 {
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if tmp != nil {
			stack = append(stack, tmp, nil)

			if tmp.Right != nil {
				stack = append(stack, tmp.Right)
			}

			if tmp.Left != nil {
				stack = append(stack, tmp.Left)
			}
		} else if len(stack) > 0 {
			ans := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, ans.Val)
		}
	}
	return res
}
