package preorder

import (
	"reflect"
	"testing"

	"github.com/peaky-lab/Daily-LeetCode/binary-tree/build"
)

func TestPreOrder(t *testing.T) {

	var bTree = build.BuildBinaryTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	resursion_res := preorder_resursion(bTree)
	stack_iteration_res := preorder_resursion(bTree)

	if !reflect.DeepEqual(resursion_res, stack_iteration_res) {
		t.Error("binary tree preorder error.")
		return
	}
	t.Log("binary tree pre order success.")
}
