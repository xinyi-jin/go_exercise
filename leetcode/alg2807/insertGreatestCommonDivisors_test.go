package alg2807

import "testing"

func Test_gcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_gcd_01", args{18, 6}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
			if got := gcdEx(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcdEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
