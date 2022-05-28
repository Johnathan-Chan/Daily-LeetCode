package inorder

import (
	"reflect"
	"testing"

	"github.com/peaky-lab/Daily-LeetCode/binary-tree/build"
)

func TestInorder(t *testing.T) {

	bTree := build.BuildBinaryTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	resursion_res := inorder_recursion(bTree)
	stack_iteration_res := inorder_stack_iteration(bTree)
	if !reflect.DeepEqual(resursion_res, stack_iteration_res) {
		t.Error("binary tree inprder error.")
		return
	}
}
