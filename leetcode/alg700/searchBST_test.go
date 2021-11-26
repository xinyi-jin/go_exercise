package alg700

import (
	"reflect"
	"testing"
)

func Test_searchBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	root := &TreeNode{Val: 4}
	l := &TreeNode{Val: 2}
	r := &TreeNode{Val: 7}
	ll := &TreeNode{Val: 1}
	rr := &TreeNode{Val: 3}
	root.Left = l
	root.Right = r
	root.Left.Left = ll
	root.Left.Right = rr

	res := &TreeNode{Val: 2}
	el := &TreeNode{Val: 1}
	er := &TreeNode{Val: 3}
	res.Left = el
	res.Right = er
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"Test1", args{root, 2}, res},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			converse(tt.args.root)
			if got := searchBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%q. searchBST() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
