package alg268

import "testing"

func Test_missingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{[]int{3, 0, 1}}, 2},
		{"Test2", args{[]int{0, 1}}, 2},
		{"Test3", args{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, 8},
		{"Test4", args{[]int{0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumberExti(tt.args.nums); got != tt.want {
				t.Errorf("%q. missingNumber() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
func BenchmarkMissingNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	}
}
func BenchmarkMissingNumberEx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		missingNumberEx([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	}
}
func BenchmarkMissingNumberExt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		missingNumberExt([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	}
}
func BenchmarkMissingNumberExti(b *testing.B) {
	for i := 0; i < b.N; i++ {
		missingNumberExti([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	}
}
