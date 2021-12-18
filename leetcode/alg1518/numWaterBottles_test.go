package alg1518

import "testing"

func Test_numWaterBottles(t *testing.T) {
	type args struct {
		numBottles  int
		numExchange int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_numWaterBottles_01", args{9, 3}, 13},
		{"Test_numWaterBottles_02", args{15, 4}, 19},
		{"Test_numWaterBottles_03", args{5, 5}, 6},
		{"Test_numWaterBottles_04", args{2, 3}, 2},
		// 借酒情况
		// {"Test_numWaterBottles_05", args{6, 3}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWaterBottles(tt.args.numBottles, tt.args.numExchange); got != tt.want {
				t.Errorf("%q. numWaterBottles() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
