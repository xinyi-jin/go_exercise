package alg807

import "testing"

func Test_maxIncreaseKeepingSkyline(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_maxIncreaseKeepingSkyline_01", args{[][]int{
			{3, 0, 8, 4},
			{2, 4, 5, 7},
			{9, 2, 6, 3},
			{0, 3, 1, 0},
		}}, 35},
		{"Test_maxIncreaseKeepingSkyline_02", args{[][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxIncreaseKeepingSkyline(tt.args.grid); got != tt.want {
				t.Errorf("%q. maxIncreaseKeepingSkyline() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
