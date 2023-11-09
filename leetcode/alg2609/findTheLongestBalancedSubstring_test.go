package alg2609

import "testing"

func Test_findTheLongestBalancedSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_findTheLongestBalancedSubstring_01", args{"01000111"}, 6},
		{"Test_findTheLongestBalancedSubstring_02", args{"00111"}, 4},
		{"Test_findTheLongestBalancedSubstring_03", args{"111"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheLongestBalancedSubstring(tt.args.s); got != tt.want {
				t.Errorf("findTheLongestBalancedSubstring() = %v, want %v", got, tt.want)
			}
			if got := findTheLongestBalancedSubstringEx(tt.args.s); got != tt.want {
				t.Errorf("findTheLongestBalancedSubstringEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
