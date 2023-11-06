package alg318

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_maxProduct_01", args{[]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}}, 16},
		{"Test_maxProduct_02", args{[]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}}, 4},
		{"Test_maxProduct_03", args{[]string{"a", "aa", "aaa", "aaaa"}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.words); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
			if got := maxProductEx(tt.args.words); got != tt.want {
				t.Errorf("maxProductEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
