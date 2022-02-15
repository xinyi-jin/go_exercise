package alg540

import "testing"

func Test_singleNonDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_singleNonDuplicate_01", args{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}}, 2},
		{"Test_singleNonDuplicate_02", args{[]int{3, 3, 7, 7, 10, 11, 11}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNonDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("singleNonDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
