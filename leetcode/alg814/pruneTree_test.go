package alg814

import (
	"reflect"
	"testing"
)

func Test_pruneTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"Test_pruneTree_01", args{&TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		}}, &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 0,
				Right: &TreeNode{
					Val: 1,
				},
			},
		}},
		{"Test_pruneTree_02", args{&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		}}, &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
				},
			},
		}},
		{"Test_pruneTree_03", args{&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		}}, &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 0,
				Right: &TreeNode{
					Val: 1,
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pruneTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pruneTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
