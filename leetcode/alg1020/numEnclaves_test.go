package alg1020

import "testing"

func Test_numEnclaves(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{"Test_numEnclaves_01", args{[][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}}, 3},
		{"Test_numEnclaves_02", args{[][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numEnclaves(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("numEnclaves() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
