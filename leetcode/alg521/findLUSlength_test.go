package alg521

import "testing"

func Test_findLUSlength(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_findLUSlength_01", args{"aba", "cdc"}, 3},
		{"Test_findLUSlength_02", args{"aaa", "bbb"}, 3},
		{"Test_findLUSlength_03", args{"aaa", "aaa"}, -1},
		{"Test_findLUSlength_04", args{"aaaa", "aaa"}, 4},
		{"Test_findLUSlength_05", args{"sadgag", "aaa"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLUSlength(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("findLUSlength() = %v, want %v", got, tt.want)
			}
		})
	}
}
