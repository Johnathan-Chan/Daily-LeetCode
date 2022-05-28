package preorder

import "github.com/peaky-lab/Daily-LeetCode/binary-tree/build"

/*
leetcode 144. 二叉树的前序遍历 https://leetcode.cn/problems/binary-tree-preorder-traversal/
binary tree preorder
recursion and stack iteration
*/

// recursion
func dfs(root *build.TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	dfs(root.Left, res)
	dfs(root.Right, res)
}

func preorder_resursion(root *build.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	dfs(root, &res)
	return res
}

//stack iteration

func preorder_stack_iteration(root *build.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	stack := []*build.TreeNode{root}

	for len(stack) > 0 {
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if tmp != nil {
			if tmp.Right != nil {
				stack = append(stack, tmp.Right)
			}

			if tmp.Left != nil {
				stack = append(stack, tmp.Left)
			}

			stack = append(stack, tmp, nil)
		} else if len(stack) > 0 {
			ans := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, ans.Val)
		}
	}
	return res
}
