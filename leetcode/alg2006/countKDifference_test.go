package alg2006

import (
	"testing"
)

func Test_countKDifference(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_countKDifference_01", args{[]int{1, 2, 2, 1}, 1}, 4},
		{"Test_countKDifference_02", args{[]int{1, 3}, 3}, 0},
		{"Test_countKDifference_03", args{[]int{3, 2, 1, 5, 4}, 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countKDifference(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countKDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countKDifferenceEx(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{"Test_countKDifferenceEx_01", args{[]int{1, 2, 2, 1}, 1}, 4},
		{"Test_countKDifferenceEx_02", args{[]int{1, 3}, 3}, 0},
		{"Test_countKDifferenceEx_03", args{[]int{3, 2, 1, 5, 4}, 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countKDifferenceEx(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("countKDifferenceEx() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func Benchmark_countKDifferenceEx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countKDifferenceEx([]int{1, 2, 2, 1}, 1)
	}
}
func Benchmark_countKDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countKDifference([]int{1, 2, 2, 1}, 1)
	}
}
