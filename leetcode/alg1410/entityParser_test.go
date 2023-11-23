package alg1410

import "testing"

func Test_entityParser(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_entityParser_01", args{"&amp; is an HTML entity but &ambassador; is not."}, "& is an HTML entity but &ambassador; is not."},
		{"Test_entityParser_02", args{"and I quote: &quot;...&quot;"}, "and I quote: \"...\""},
		{"Test_entityParser_03", args{"Stay home! Practice on Leetcode :)"}, "Stay home! Practice on Leetcode :)"},
		{"Test_entityParser_04", args{"x &gt; y &amp;&amp; x &lt; y is always false"}, "x > y && x < y is always false"},
		{"Test_entityParser_05", args{"leetcode.com&frasl;problemset&frasl;all"}, "leetcode.com/problemset/all"},
		{"Test_entityParser_06", args{"&amp;gt;"}, "&gt;"},
		{"Test_entityParser_07", args{"&&gt;"}, "&>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := entityParser(tt.args.text); got != tt.want {
				t.Errorf("entityParser() = %v, want %v", got, tt.want)
			}
			if got := entityParserEx(tt.args.text); got != tt.want {
				t.Errorf("entityParserEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
