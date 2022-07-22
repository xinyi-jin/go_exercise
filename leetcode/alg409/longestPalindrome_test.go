package alg409

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_longestPalindrome_01", args{"abccccdd"}, 7},
		{"Test_longestPalindrome_02", args{"a"}, 1},
		{"Test_longestPalindrome_03", args{"bb"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
			if got := longestPalindromeEx(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
