package alg495

import (
	"testing"
)

func Test_findPoisonedDuration(t *testing.T) {
	type args struct {
		timeSeries []int
		duration   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{[]int{1, 4}, 2}, 4},
		{"Test2", args{[]int{1, 2}, 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPoisonedDuration(tt.args.timeSeries, tt.args.duration); got != tt.want {
				t.Errorf("%q. findPoisonedDuration() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Benchmark_FindPoisonedDuration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findPoisonedDuration([]int{1, 4}, 2)
	}
}
func Benchmark_FindPoisonedDurationEx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findPoisonedDurationEx([]int{1, 4}, 2)
	}
}
