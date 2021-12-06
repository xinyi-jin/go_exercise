package alg1816

import "testing"

func Test_truncateSentence(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_truncateSentence_01", args{"Hello how are you Contestant", 4}, "Hello how are you"},
		{"Test_truncateSentence_02", args{"What is the solution to this problem", 4}, "What is the solution"},
		{"Test_truncateSentence_03", args{"chopper is not a tanuki", 5}, "chopper is not a tanuki"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := truncateSentence(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("%q. truncateSentence() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
func Benchmark_truncateSentence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		truncateSentence("What is the solution to this problem", 4)
	}
}

func Benchmark_truncateSentenceEx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		truncateSentenceEx("What is the solution to this problem", 4)
	}
}
