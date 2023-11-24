package alg2824

import "testing"

func Test_countPairs(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_countPairs_01", args{[]int{-1, 1, 2, 3, 1}, 2}, 3},
		{"Test_countPairs_02", args{[]int{-6, 2, 5, -2, -7, -1, 3}, -2}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
