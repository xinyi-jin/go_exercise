package alg2073

import "testing"

func Test_timeRequiredToBuy(t *testing.T) {
	type args struct {
		tickets []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_timeRequiredToBuy_01", args{[]int{2, 3, 2}, 2}, 6},
		{"Test_timeRequiredToBuy_02", args{[]int{5, 1, 1, 1}, 0}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeRequiredToBuy(tt.args.tickets, tt.args.k); got != tt.want {
				t.Errorf("timeRequiredToBuy() = %v, want %v", got, tt.want)
			}
		})
	}
}
