package postorder

import (
	"reflect"
	"testing"

	"github.com/peaky-lab/Daily-LeetCode/binary-tree/build"
)

func TestPostorder(t *testing.T) {
	bTree := build.BuildBinaryTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	resursion_res := postorder_recursion(bTree)
	stack_iteration_res := postorder_stack_iteration(bTree)
	if !reflect.DeepEqual(resursion_res, stack_iteration_res) {
		t.Error("binary tree postorder error.")
		return
	}
}
