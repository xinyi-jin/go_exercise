package alg2441

import (
	"testing"
)

func Test_findMaxK(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_findMaxK_01", args{[]int{-1, 2, -3, 3}}, 3},
		{"Test_findMaxK_02", args{[]int{-1, 10, 6, 7, -7, 1}}, 7},
		{"Test_findMaxK_03", args{[]int{-10, 8, 6, 7, -2, -3}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxK(tt.args.nums); got != tt.want {
				t.Errorf("findMaxK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxKEx(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_findMaxKEx_01", args{[]int{-1, 2, -3, 3}}, 3},
		{"Test_findMaxKEx_02", args{[]int{-1, 10, 6, 7, -7, 1}}, 7},
		{"Test_findMaxKEx_03", args{[]int{-10, 8, 6, 7, -2, -3}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxKEx(tt.args.nums); got != tt.want {
				t.Errorf("findMaxKEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
