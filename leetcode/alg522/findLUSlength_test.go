package alg522

import "testing"

func Test_findLUSlength(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_findLUSlength_01", args{[]string{"aba", "cdc", "eae"}}, 3},
		{"Test_findLUSlength_02", args{[]string{"aaa", "aaa", "aa"}}, -1},
		// {"Test_findLUSlength_03", args{[]string{"aba", "cdc", "eae"}}, 3},
		{"Test_findLUSlength_04", args{[]string{"aabbcc", "aabbcc", "e"}}, 1},
		{"Test_findLUSlength_05", args{[]string{"aaa", "abc"}}, 3},
		{"Test_findLUSlength_06", args{[]string{"aabbcc", "aabbcc", "cb", "abc"}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLUSlength(tt.args.strs); got != tt.want {
				t.Errorf("findLUSlength() = %v, want %v", got, tt.want)
			}
		})
	}
}
