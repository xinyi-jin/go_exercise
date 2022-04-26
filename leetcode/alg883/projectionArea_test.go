package alg883

import "testing"

func Test_projectionArea(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_projectionArea_01", args{[][]int{{1, 2}, {3, 4}}}, 17},
		{"Test_projectionArea_02", args{[][]int{{2}}}, 5},
		{"Test_projectionArea_03", args{[][]int{{1, 0}, {0, 2}}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := projectionArea(tt.args.grid); got != tt.want {
				t.Errorf("projectionArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
