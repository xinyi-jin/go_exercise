package alg2656

import "testing"

func Test_maximizeSum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_maximizeSum_01", args{[]int{1, 2, 3, 4, 5}, 3}, 18},
		{"Test_maximizeSum_02", args{[]int{5, 5, 5}, 2}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximizeSum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximizeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
