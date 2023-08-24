package alg1267

import "testing"

func Test_countServers(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_countServers_01", args{[][]int{
			{1, 0},
			{0, 1},
		}}, 0},
		{"Test_countServers_02", args{[][]int{
			{1, 0},
			{1, 1},
		}}, 3},
		{"Test_countServers_03", args{[][]int{
			{1, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 1, 0},
		}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countServers(tt.args.grid); got != tt.want {
				t.Errorf("countServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
