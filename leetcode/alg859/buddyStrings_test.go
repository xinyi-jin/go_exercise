package alg859

import "testing"

func Test_buddyStrings(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test1", args{"ab", "ba"}, true},
		{"Test2", args{"ab", "ab"}, false},
		{"Test3", args{"aa", "aa"}, true},
		{"Test4", args{"aaaaaaabc", "aaaaaaacb"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("%q. buddyStrings() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
