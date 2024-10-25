package alg12

import (
	"testing"
)

func TestIntToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{3}, "III"},
		{"2", args{4}, "IV"},
		{"3", args{9}, "IX"},
		{"4", args{58}, "LVIII"},
		{"5", args{1994}, "MCMXCIV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToRoman(tt.args.num); got != tt.want {
				t.Errorf("IntToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToRomanEx(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{3}, "III"},
		{"2", args{4}, "IV"},
		{"3", args{9}, "IX"},
		{"4", args{58}, "LVIII"},
		{"5", args{1994}, "MCMXCIV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToRomanEx(tt.args.num); got != tt.want {
				t.Errorf("IntToRomanEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{3}, "III"},
		{"2", args{4}, "IV"},
		{"3", args{9}, "IX"},
		{"4", args{58}, "LVIII"},
		{"5", args{1994}, "MCMXCIV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
