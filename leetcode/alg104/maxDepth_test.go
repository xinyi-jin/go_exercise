package alg104

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_maxDepth_01", args{&TreeNode{Val: 10}}, 1},
		{"Test_maxDepth_02", args{&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 5,
				},
			},
		}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_maxDepthEx(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_maxDepthEx_01", args{&TreeNode{Val: 10}}, 1},
		{"Test_maxDepthEx_02", args{&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 5,
				},
			},
		}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepthEx(tt.args.root); got != tt.want {
				t.Errorf("maxDepthEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
