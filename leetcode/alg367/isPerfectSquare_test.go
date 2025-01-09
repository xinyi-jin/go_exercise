package alg367

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test_isPerfectSquare_01", args{16}, true},
		{"Test_isPerfectSquare_02", args{14}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
			if got := isPerfectSquare2(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare2() = %v, want %v", got, tt.want)
			}
		})
	}
}
