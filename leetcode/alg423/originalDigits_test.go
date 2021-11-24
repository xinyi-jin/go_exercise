package alg423

import (
	"reflect"
	"testing"
)

func Test_getNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want map[byte]int
	}{
		{"Test1", args{"zeroonetwothreefourfivesixseveneightnine"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumber(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%q. getNumber() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_originalDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_originalDigits_1", args{"owoztneoer"}, "012"},
		{"Test_originalDigits_2", args{"fviefuro"}, "45"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := originalDigits(tt.args.s); got != tt.want {
				t.Errorf("%q. originalDigits() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
