package alg1749

import "testing"

func Test_maxAbsoluteSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_maxAbsoluteSum_01", args{[]int{1, -3, 2, 3, -4}}, 5},
		{"Test_maxAbsoluteSum_02", args{[]int{2, -5, 1, -4, 3, -2}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAbsoluteSum(tt.args.nums); got != tt.want {
				t.Errorf("maxAbsoluteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
