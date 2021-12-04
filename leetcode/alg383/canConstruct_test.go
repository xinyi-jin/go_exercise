package alg383

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test_canConstruct_1", args{"a", "b"}, false},
		{"Test_canConstruct_2", args{"aa", "ab"}, false},
		{"Test_canConstruct_3", args{"aa", "aab"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("%q. canConstruct() = %v, want %v", tt.name, got, tt.want)
			}
			if got := canConstructEx(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("%q. canConstructEx() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
