package alg1038

import (
	"testing"
)

func Test_bstToGst(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"Test_bstToGst_01", args{&TreeNode{
			4,
			&TreeNode{
				1,
				&TreeNode{
					0, nil, nil,
				},
				&TreeNode{
					2,
					nil,
					&TreeNode{
						3, nil, nil,
					},
				},
			},
			&TreeNode{
				6,
				&TreeNode{
					5, nil, nil,
				},
				&TreeNode{
					7, nil,
					&TreeNode{
						8, nil, nil,
					},
				},
			},
		}}, &TreeNode{
			30,
			&TreeNode{
				36,
				&TreeNode{
					36, nil, nil,
				},
				&TreeNode{
					35,
					nil,
					&TreeNode{
						33, nil, nil,
					},
				},
			},
			&TreeNode{
				21,
				&TreeNode{
					26, nil, nil,
				},
				&TreeNode{
					15, nil,
					&TreeNode{
						8, nil, nil,
					},
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			/* if got := bstToGst(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstToGst() = %v, want %v", got, tt.want)
			} */
			if got := bstToGst(tt.args.root); !TreeEqual(got, tt.want) {
				t.Errorf("bstToGst() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 两树各个节点值是否相等
func TreeEqual(r1, r2 *TreeNode) bool {
	var dfs func(a, b *TreeNode) bool

	dfs = func(a, b *TreeNode) bool {
		if a != nil && b == nil || a == nil && b != nil {
			return false
		}
		if a != nil && b != nil {
			res := dfs(a.Left, b.Left)
			if !res {
				return res
			}
			if a.Val != b.Val {
				res = false
				return res
			}
			res = dfs(a.Right, b.Right)
			if !res {
				return res
			}
		}
		return true
	}

	return dfs(r1, r2)
}
