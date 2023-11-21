package alg2216

import "testing"

func Test_minDeletion(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_minDeletion_01", args{[]int{1, 1, 2, 3, 5}}, 1},
		{"Test_minDeletion_02", args{[]int{1, 1, 2, 2, 3, 3}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDeletion(tt.args.nums); got != tt.want {
				t.Errorf("minDeletion() = %v, want %v", got, tt.want)
			}
		})
	}
}
