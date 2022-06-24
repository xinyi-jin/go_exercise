package alg515

import (
	"reflect"
	"testing"
)

func Test_largestValues(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test_01", args{
			&TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 3,
					},
					Val: 3,
				},
				Right: &TreeNode{
					Right: &TreeNode{
						Val: 9,
					},
					Val: 2,
				},
				Val: 1,
			},
		}, []int{1, 3, 9}},
		{"Test_02", args{
			&TreeNode{
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
				Val: 1,
			},
		}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestValues(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestValues() = %v, want %v", got, tt.want)
			}
			if got := largestValuesEx(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestValuesEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_largestValues(b *testing.B) {
	for i := 0; i < b.N; i++ {
		largestValues(&TreeNode{
			Left: &TreeNode{
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 3,
				},
				Val: 3,
			},
			Right: &TreeNode{
				Right: &TreeNode{
					Val: 9,
				},
				Val: 2,
			},
			Val: 1,
		})
	}
}

func Benchmark_largestValuesEx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		largestValuesEx(&TreeNode{
			Left: &TreeNode{
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 3,
				},
				Val: 3,
			},
			Right: &TreeNode{
				Right: &TreeNode{
					Val: 9,
				},
				Val: 2,
			},
			Val: 1,
		})
	}
}
