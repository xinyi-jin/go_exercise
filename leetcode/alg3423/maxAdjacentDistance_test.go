package alg3423

import "testing"

func Test_maxAdjacentDistance(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_maxAdjacentDistance_01", args{[]int{1, 2, 4}}, 3},
		{"Test_maxAdjacentDistance_02", args{[]int{-5, -10, -5}}, 5},
		{"Test_maxAdjacentDistance_03", args{[]int{-4, -2, -3}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAdjacentDistance(tt.args.nums); got != tt.want {
				t.Errorf("maxAdjacentDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}
