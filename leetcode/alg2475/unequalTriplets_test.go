package alg2475

import "testing"

func Test_unequalTriplets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_unequalTriplets_01", args{[]int{4, 4, 2, 4, 3}}, 3},
		{"Test_unequalTriplets_02", args{[]int{1, 1, 1, 1, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unequalTriplets(tt.args.nums); got != tt.want {
				t.Errorf("unequalTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
