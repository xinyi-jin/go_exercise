package alg2477

import "testing"

func Test_minimumFuelCost(t *testing.T) {
	type args struct {
		roads [][]int
		seats int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Test_minimumFuelCost_01", args{[][]int{
			{0, 1},
			{0, 2},
			{0, 3},
		}, 5}, 3},
		{"Test_minimumFuelCost_02", args{[][]int{
			{3, 1},
			{3, 2},
			{1, 0},
			{0, 4},
			{0, 5},
			{4, 6},
		}, 2}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumFuelCost(tt.args.roads, tt.args.seats); got != tt.want {
				t.Errorf("minimumFuelCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
