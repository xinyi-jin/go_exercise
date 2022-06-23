package alg30

import (
	"reflect"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test_01", args{"barfoothefoobarman", []string{"foo", "bar"}}, []int{0, 9}},
		{"Test_02", args{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}}, []int{}},
		{"Test_03", args{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}}, []int{6, 9, 12}},
		{"Test_04", args{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}}, []int{8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_findSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"})
	}
}
func Benchmark_findSubstringEx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findSubstringEx("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"})
	}
}
