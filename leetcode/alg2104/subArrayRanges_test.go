package alg2104

import "testing"

func Test_subArrayRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Test_subArrayRanges_01", args{[]int{1, 2, 3}}, 4},
		{"Test_subArrayRanges_02", args{[]int{1, 3, 3}}, 4},
		{"Test_subArrayRanges_03", args{[]int{4, -2, -3, 4, 1}}, 59},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subArrayRanges(tt.args.nums); got != tt.want {
				t.Errorf("subArrayRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_subArrayRangesEx(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// {"Test_subArrayRangesEx_01", args{[]int{1, 2, 3}}, 4},
		// {"Test_subArrayRangesEx_02", args{[]int{1, 3, 3}}, 4},
		{"Test_subArrayRangesEx_03", args{[]int{4, -2, -3, 4, 1}}, 59},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subArrayRangesEx(tt.args.nums); got != tt.want {
				t.Errorf("subArrayRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_subArrayRangesExtion(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Test_subArrayRangesExtion_01", args{[]int{1, 2, 3}}, 4},
		{"Test_subArrayRangesExtion_02", args{[]int{1, 3, 3}}, 4},
		{"Test_subArrayRangesExtion_03", args{[]int{4, -2, -3, 4, 1}}, 59},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subArrayRangesExtion(tt.args.nums); got != tt.want {
				t.Errorf("subArrayRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
