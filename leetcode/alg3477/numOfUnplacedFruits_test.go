package alg3477

import "testing"

func Test_numOfUnplacedFruits(t *testing.T) {
	type args struct {
		fruits  []int
		baskets []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_numOfUnplacedFruits_01", args{[]int{4, 2, 5}, []int{3, 5, 4}}, 1},
		{"Test_numOfUnplacedFruits_02", args{[]int{3, 6, 1}, []int{6, 4, 7}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfUnplacedFruits(tt.args.fruits, tt.args.baskets); got != tt.want {
				t.Errorf("numOfUnplacedFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
