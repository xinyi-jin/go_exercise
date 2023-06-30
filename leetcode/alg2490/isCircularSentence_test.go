package alg2490

import "testing"

func Test_isCircularSentence(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test_isCircularSentence_01", args{"leetcode exercises sound delightful"}, true},
		{"Test_isCircularSentence_02", args{"eetcode"}, true},
		{"Test_isCircularSentence_03", args{"Leetcode is cool"}, false},
		{"Test_isCircularSentence_04", args{"a b"}, false},
		{"Test_isCircularSentence_05", args{"a a"}, true},
		{"Test_isCircularSentence_06", args{"a"}, true},
		{"Test_isCircularSentence_07", args{"IuTiUtGGsNydmacGduehPPGksKQyT TmOraUbCcQdnZUCpGCYtGp p pG GCcRvZDRawqGKOiBSLwjIDOjdhnHiisfddYoeHqxOqkUvOEyI"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCircularSentence(tt.args.sentence); got != tt.want {
				t.Errorf("isCircularSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
