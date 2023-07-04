package alg2679

import "testing"

func Test_matrixSum(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_matrixSum_01", args{[][]int{
			{7, 2, 1},
			{6, 4, 2},
			{6, 5, 3},
			{3, 2, 1},
		}}, 15},
		{"Test_matrixSum_02", args{[][]int{
			{1},
		}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixSum(tt.args.nums); got != tt.want {
				t.Errorf("matrixSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
