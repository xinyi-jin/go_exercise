package alg653

import (
	"reflect"
	"testing"
)

func Test_findTarget(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test_findTarget_01", args{root: convertTree([]int{5, 3, 6, 2, 4, 0, 7}), k: 9}, true},
		{"Test_findTarget_02", args{root: convertTree([]int{5, 3, 6, 2, 4, 0, 7}), k: 28}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTarget(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertTreeNode(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"Test_convertTreeNode_01", args{[]int{4, 2, 6, 1, 3, 5, 7}}, convertTree([]int{4, 2, 6, 1, 3, 5, 7})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				traverTree(got)
				t.Errorf("convertTreeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
