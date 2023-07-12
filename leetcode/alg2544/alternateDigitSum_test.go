package alg2544

import "testing"

func Test_alternateDigitSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_alternateDigitSum_01", args{521}, 4},
		{"Test_alternateDigitSum_02", args{111}, 1},
		{"Test_alternateDigitSum_03", args{886996}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alternateDigitSum(tt.args.n); got != tt.want {
				t.Errorf("alternateDigitSum() = %v, want %v", got, tt.want)
			}
			if got := alternateDigitSumEx(tt.args.n); got != tt.want {
				t.Errorf("alternateDigitSumEx() = %v, want %v", got, tt.want)
			}
			if got := alternateDigitSumExtion(tt.args.n); got != tt.want {
				t.Errorf("alternateDigitSumExtion() = %v, want %v", got, tt.want)
			}
		})
	}
}
