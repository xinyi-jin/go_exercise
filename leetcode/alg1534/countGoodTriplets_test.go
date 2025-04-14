package alg1534

import "testing"

func Test_countGoodTriplets(t *testing.T) {
	type args struct {
		arr []int
		a   int
		b   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_countGoodTriplets_01", args{[]int{3, 0, 1, 1, 9, 7}, 7, 2, 3}, 4},
		{"Test_countGoodTriplets_02", args{[]int{1, 1, 2, 2, 3}, 0, 0, 1}, 0},
		{"Test_countGoodTriplets_03", args{[]int{7, 3, 7, 3, 12, 1, 12, 2, 3}, 5, 8, 1}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodTriplets(tt.args.arr, tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("countGoodTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
