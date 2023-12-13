package alg2697

import "testing"

func Test_makeSmallestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_makeSmallestPalindrome_01", args{"egcfe"}, "efcfe"},
		{"Test_makeSmallestPalindrome_02", args{"abcd"}, "abba"},
		{"Test_makeSmallestPalindrome_03", args{"seven"}, "neven"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeSmallestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("makeSmallestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
