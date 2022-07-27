package alg592

import (
	"testing"
)

func Test_fractionAddition(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_fractionAddition_01", args{"-1/2+1/2"}, "0/1"},
		{"Test_fractionAddition_02", args{"-1/2+1/2+1/3"}, "1/3"},
		{"Test_fractionAddition_03", args{"1/3-1/2"}, "-1/6"},
		{"Test_fractionAddition_04", args{"-5/2+10/3+7/9"}, "29/18"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fractionAddition(tt.args.expression); got != tt.want {
				t.Errorf("fractionAddition() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.x); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
