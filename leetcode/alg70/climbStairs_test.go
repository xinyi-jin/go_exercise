package alg70

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_climbStairs_01", args{2}, 2},
		{"Test_climbStairs_02", args{3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
			if got := climbStairsEx1(tt.args.n); got != tt.want {
				t.Errorf("climbStairsEx1() = %v, want %v", got, tt.want)
			}
			if got := climbStairsEx2(tt.args.n); got != tt.want {
				t.Errorf("climbStairsEx2() = %v, want %v", got, tt.want)
			}
			if got := climbStairsEx3(tt.args.n); got != tt.want {
				t.Errorf("climbStairsEx3() = %v, want %v", got, tt.want)
			}
		})
	}
}
