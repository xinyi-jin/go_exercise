package alg1072

import "testing"

func Test_maxEqualRowsAfterFlips(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_maxEqualRowsAfterFlips_01", args{
			matrix: [][]int{
				{0, 1},
				{1, 1},
			},
		}, 1},
		{"Test_maxEqualRowsAfterFlips_02", args{
			matrix: [][]int{
				{0, 1},
				{1, 0},
			},
		}, 2},
		{"Test_maxEqualRowsAfterFlips_03", args{
			matrix: [][]int{
				{0, 0, 0},
				{0, 0, 1},
				{1, 1, 0},
			},
		}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEqualRowsAfterFlips(tt.args.matrix); got != tt.want {
				t.Errorf("maxEqualRowsAfterFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
