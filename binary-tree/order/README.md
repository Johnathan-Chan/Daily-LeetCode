# binary tree order
- [binary tree order](#binary-tree-order)
	- [order template](#order-template)
		- [recursion](#recursion)
		- [stack iterattion](#stack-iterattion)
	- [levelorder template](#levelorder-template)


## order template
### recursion
```golang
/*
Example bTree preorder
*/

func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	dfs(root.Left, res)
	dfs(root.Right, res)
}

func preorder_resursion(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	dfs(root, &res)
	return res
}
```
### stack iterattion
```golang
/*
Example bTree preorder
*/

func preorder_stack_iteration(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	stack := []*TreeNode{root}

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
```

## levelorder template
```golang
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*TreeNode{root}
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
```