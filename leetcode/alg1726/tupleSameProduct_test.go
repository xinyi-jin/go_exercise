package alg1726

import "testing"

func Test_tupleSameProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_tupleSameProduct_01", args{[]int{2, 3, 4, 6}}, 8},
		{"Test_tupleSameProduct_02", args{[]int{1, 2, 4, 5, 10}}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tupleSameProduct(tt.args.nums); got != tt.want {
				t.Errorf("tupleSameProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
