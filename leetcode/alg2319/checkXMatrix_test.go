package alg2319

import "testing"

func Test_checkXMatrix(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test_checkXMatrix_01", args{[][]int{
			{5, 7, 0},
			{0, 3, 1},
			{0, 5, 0},
		}}, false},
		{"Test_checkXMatrix_02", args{[][]int{
			{2, 0, 0, 1},
			{0, 3, 1, 0},
			{0, 5, 2, 0},
			{4, 0, 0, 2},
		}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkXMatrix(tt.args.grid); got != tt.want {
				t.Errorf("checkXMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
