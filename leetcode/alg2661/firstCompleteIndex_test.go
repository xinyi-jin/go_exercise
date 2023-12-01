package alg2661

import "testing"

func Test_firstCompleteIndex(t *testing.T) {
	type args struct {
		arr []int
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_firstCompleteIndex_01", args{[]int{1, 3, 4, 2}, [][]int{
			{1, 4},
			{2, 3},
		}}, 2},
		{"Test_firstCompleteIndex_02", args{[]int{2, 8, 7, 4, 1, 3, 5, 6, 9}, [][]int{
			{3, 2, 5},
			{1, 4, 6},
			{8, 7, 9},
		}}, 3},
		{"Test_firstCompleteIndex_03", args{[]int{1, 4, 5, 2, 6, 3}, [][]int{
			{4, 3, 5},
			{1, 2, 6},
		}}, 1},
		{"Test_firstCompleteIndex_04", args{[]int{1, 4, 5, 2, 6, 3}, [][]int{
			{4, 3},
			{1, 2},
			{5, 6},
		}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstCompleteIndex(tt.args.arr, tt.args.mat); got != tt.want {
				t.Errorf("firstCompleteIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
