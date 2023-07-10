package alg16

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_threeSumClosest_01", args{[]int{-1, 2, 1, -4}, 1}, 2},
		{"Test_threeSumClosest_02", args{[]int{0, 0, 0}, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
			if got := threeSumClosestEx(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosestEx() = %v, want %v", got, tt.want)
			}
			if got := threeSumClosestExtion(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosestExtion() = %v, want %v", got, tt.want)
			}
		})
	}
}
