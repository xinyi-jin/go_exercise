package alg108

import (
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test_01", args{[]int{-10, -3, 0, 5, 9}}},
		{"Test_02", args{[]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedArrayToBST(tt.args.nums)
			t.Logf("sortedArrayToBST(%v)", t.Name())
			traverse(got)
		})
	}
}
