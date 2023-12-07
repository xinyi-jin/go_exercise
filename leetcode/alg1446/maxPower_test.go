package alg1446

import "testing"

func Test_maxPower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_maxPower_1", args{"leetcode"}, 2},
		{"Test_maxPower_2", args{"abbcccddddeeeeedcba"}, 5},
		{"Test_maxPower_3", args{"triplepillooooow"}, 5},
		{"Test_maxPower_4", args{"hooraaaaaaaaaaay"}, 11},
		{"Test_maxPower_5", args{"tourist"}, 1},
		{"Test_maxPower_6", args{"a"}, 1},
		{"Test_maxPower_7", args{"cc"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPower(tt.args.s); got != tt.want {
				t.Errorf("%q. maxPower() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
