package alg1657

import "testing"

func Test_closeStrings(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test_closeStrings_01", args{"abc", "bca"}, true},
		{"Test_closeStrings_02", args{"a", "aa"}, false},
		{"Test_closeStrings_03", args{"cabbba", "abbccc"}, true},
		{"Test_closeStrings_04", args{"cabbba", "aabbss"}, false},
		{"Test_closeStrings_05", args{"kmihff", "kffmmi"}, false},
		{"Test_closeStrings_06", args{"aabbcccddd", "abccccdddd"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closeStrings(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("closeStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
