package alg670

import "testing"

func Test_maximumSwap(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_maximumSwap_01", args{2736}, 7236},
		{"Test_maximumSwap_02", args{9973}, 9973},
		{"Test_maximumSwap_03", args{98368}, 98863},
		{"Test_maximumSwap_04", args{1993}, 9913},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSwap(tt.args.num); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
