package alg559

import (
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_maxDepth_01", args{&Node{
			Val: 1,
			Children: []*Node{
				{Val: 3},
				{Val: 2},
				{Val: 4, Children: []*Node{
					{Val: 5},
					{Val: 6},
				},
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

func Test_max(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDepthEx(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{"Test_maxDepthEx_01", args{&Node{
			Val: 1,
			Children: []*Node{
				{Val: 3},
				{Val: 2},
				{Val: 4, Children: []*Node{
					{Val: 5},
					{Val: 6},
				},
				},
			},
		}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxDepthEx(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("maxDepthEx() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
