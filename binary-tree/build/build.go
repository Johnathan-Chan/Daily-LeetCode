package build

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildBinaryTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	i := 0
	for index, value := range inorder {
		if value == preorder[0] {
			i = index
			break
		}
	}

	root.Left = BuildBinaryTree(preorder[1:i+1], inorder[:i])
	root.Right = BuildBinaryTree(preorder[i+1:], inorder[i+1:])
	return root
}
