package alg03

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{"abcabcbb"}, 3},
		{"Test3", args{"pwwkew"}, 3},
		{"Test4", args{""}, 0},
		{"Test5", args{"aa"}, 1},
		{"Test5", args{"au"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_lengthOfLongestSubstringBaoLi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test1", args{"abcabcbb"}, 3},
		{"Test2", args{"bbbbb"}, 1},
		{"Test3", args{"pwwkew"}, 3},
		{"Test4", args{""}, 0},
		{"Test5", args{"aa"}, 1},
		{"Test6", args{"au"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringBaoLi(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringBaoLi() = %v, want %v", got, tt.want)
			}
		})
	}
}
