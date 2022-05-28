package depth

import (
	"testing"

	"github.com/peaky-lab/Daily-LeetCode/binary-tree/build"
)

func TestMaxDepth(t *testing.T) {

	bTree := build.BuildBinaryTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	res := 3
	if maxDepth(bTree) != res {
		t.Error("binary tree max depth error.")
		return
	}
}

func TestMinDepth(t *testing.T) {
	bTree := build.BuildBinaryTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	res := 2
	if minDepth(bTree) != res {
		t.Error("binary tree min depth error.")
		return
	}
}
