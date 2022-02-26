package alg2016

import "testing"

func Test_maximumDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_maximumDifference_01", args{[]int{7, 1, 5, 4}}, 4},
		{"Test_maximumDifference_02", args{[]int{9, 4, 3, 2}}, -1},
		{"Test_maximumDifference_03", args{[]int{1, 5, 2, 10}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumDifference(tt.args.nums); got != tt.want {
				t.Errorf("maximumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_maximumDifferenceEx(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_maximumDifferenceEx_01", args{[]int{7, 1, 5, 4}}, 4},
		{"Test_maximumDifferenceEx_02", args{[]int{9, 4, 3, 2}}, -1},
		{"Test_maximumDifferenceEx_03", args{[]int{1, 5, 2, 10}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumDifferenceEx(tt.args.nums); got != tt.want {
				t.Errorf("maximumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
