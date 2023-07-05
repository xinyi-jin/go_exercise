package alg2600

import "testing"

func Test_kItemsWithMaximumSum(t *testing.T) {
	type args struct {
		numOnes    int
		numZeros   int
		numNegOnes int
		k          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_kItemsWithMaximumSum_01", args{3, 2, 0, 2}, 2},
		{"Test_kItemsWithMaximumSum_02", args{3, 2, 0, 4}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kItemsWithMaximumSum(tt.args.numOnes, tt.args.numZeros, tt.args.numNegOnes, tt.args.k); got != tt.want {
				t.Errorf("kItemsWithMaximumSum() = %v, want %v", got, tt.want)
			}
			if got := kItemsWithMaximumSumEx(tt.args.numOnes, tt.args.numZeros, tt.args.numNegOnes, tt.args.k); got != tt.want {
				t.Errorf("kItemsWithMaximumSumEx() = %v, want %v", got, tt.want)
			}
			if got := kItemsWithMaximumSumExtion(tt.args.numOnes, tt.args.numZeros, tt.args.numNegOnes, tt.args.k); got != tt.want {
				t.Errorf("kItemsWithMaximumSumExtion() = %v, want %v", got, tt.want)
			}
		})
	}
}
