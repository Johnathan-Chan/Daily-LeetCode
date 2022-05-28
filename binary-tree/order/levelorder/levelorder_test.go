package levelorder

import (
	"reflect"
	"testing"

	"github.com/peaky-lab/Daily-LeetCode/binary-tree/build"
)

func TestLevelorder(t *testing.T) {

	bTree := build.BuildBinaryTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	expectedRes := [][]int{{3}, {9, 20}, {15, 7}}

	if !reflect.DeepEqual(levelOrder(bTree), expectedRes) {
		t.Error("binary tree levelorder error.")
		return
	}
}
