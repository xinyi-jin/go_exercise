package alg868

import "testing"

func Test_binaryGap(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_binaryGap_01", args{22}, 2},
		{"Test_binaryGap_02", args{8}, 0},
		{"Test_binaryGap_03", args{5}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryGap(tt.args.n); got != tt.want {
				t.Errorf("binaryGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
