package alg3356

import "testing"

func Test_minZeroArray(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_minZeroArray_01", args{[]int{2, 0, 2}, [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}}}, 2},
		{"Test_minZeroArray_02", args{[]int{4, 3, 2, 1}, [][]int{{1, 3, 2}, {0, 2, 1}}}, -1},
		{"Test_minZeroArray_03", args{[]int{0}, [][]int{{0, 0, 2}, {0, 0, 4}, {0, 0, 4}, {0, 0, 3}, {0, 0, 5}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minZeroArray(tt.args.nums, tt.args.queries); got != tt.want {
				t.Errorf("minZeroArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
