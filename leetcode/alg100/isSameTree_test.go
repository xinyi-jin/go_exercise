package alg100

import "testing"

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test_01", args{
			p: &TreeNode{
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
				Val: 1,
			},
			q: &TreeNode{
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
				Val: 1,
			}}, true},
		{"Test_02", args{
			p: &TreeNode{
				Left: &TreeNode{
					Val: 2,
				},
				Val: 1,
			},
			q: &TreeNode{
				Right: &TreeNode{
					Val: 2,
				},
				Val: 1,
			}}, false},
		{"Test_03", args{
			p: &TreeNode{
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 1,
				},
				Val: 1,
			},
			q: &TreeNode{
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
				Val: 1,
			}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
			if got := isSameTreeEx(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTreeEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
