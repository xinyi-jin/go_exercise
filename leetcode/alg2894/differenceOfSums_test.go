package alg2894

import "testing"

func Test_differenceOfSums(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_differenceOfSums_01", args{10, 3}, 19},
		{"Test_differenceOfSums_02", args{5, 6}, 15},
		{"Test_differenceOfSums_03", args{5, 1}, -15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := differenceOfSums(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("differenceOfSums() = %v, want %v", got, tt.want)
			}
			if got := differenceOfSumsEx(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("differenceOfSumsEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
