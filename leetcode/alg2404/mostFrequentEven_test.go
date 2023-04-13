package alg2404

import "testing"

func Test_mostFrequentEven(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test_01", args{[]int{0, 1, 2, 2, 4, 4, 1}}, 2},
		{"test_02", args{[]int{4, 4, 4, 9, 2, 4}}, 4},
		{"test_03", args{[]int{29, 47, 21, 41, 13, 37, 25, 7}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostFrequentEven(tt.args.nums); got != tt.want {
				t.Errorf("mostFrequentEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
