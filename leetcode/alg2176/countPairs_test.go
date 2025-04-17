package alg2176

import "testing"

func Test_countPairs(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_countPairs_01", args{[]int{3, 1, 2, 2, 2, 1, 3}, 2}, 4},
		{"Test_countPairs_02", args{[]int{1, 2, 3, 4}, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
