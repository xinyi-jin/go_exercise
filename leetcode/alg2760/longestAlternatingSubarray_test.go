package alg2760

import "testing"

func Test_longestAlternatingSubarray(t *testing.T) {
	type args struct {
		nums      []int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_longestAlternatingSubarray_01", args{[]int{3, 2, 5, 4}, 5}, 3},
		{"Test_longestAlternatingSubarray_02", args{[]int{1, 2}, 2}, 1},
		{"Test_longestAlternatingSubarray_03", args{[]int{2, 3, 4, 5}, 4}, 3},
		{"Test_longestAlternatingSubarray_04", args{[]int{4, 10, 3}, 10}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestAlternatingSubarray(tt.args.nums, tt.args.threshold); got != tt.want {
				t.Errorf("longestAlternatingSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
