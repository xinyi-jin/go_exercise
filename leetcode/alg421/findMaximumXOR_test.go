package alg421

import "testing"

func Test_findMaximumXOR(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_findMaximumXOR_01", args{[]int{3, 10, 5, 25, 2, 8}}, 28},
		{"Test_findMaximumXOR_02", args{[]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}}, 127},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximumXOR(tt.args.nums); got != tt.want {
				t.Errorf("findMaximumXOR() = %v, want %v", got, tt.want)
			}
			if got := findMaximumXOREx(tt.args.nums); got != tt.want {
				t.Errorf("findMaximumXOREx() = %v, want %v", got, tt.want)
			}
			if got := findMaximumXORExtion(tt.args.nums); got != tt.want {
				t.Errorf("findMaximumXORExtion() = %v, want %v", got, tt.want)
			}
		})
	}
}
