package alg2586

import "testing"

func Test_vowelStrings(t *testing.T) {
	type args struct {
		words []string
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_vowelStrings_01", args{[]string{"are", "amy", "u"}, 0, 2}, 2},
		{"Test_vowelStrings_02", args{[]string{"hey", "aeo", "mu", "ooo", "artro"}, 1, 4}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vowelStrings(tt.args.words, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("vowelStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
