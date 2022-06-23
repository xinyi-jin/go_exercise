package alg242

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test_01", args{"anagram", "nagaram"}, true},
		{"Test_02", args{"rat", "car"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
			if got := isAnagramEx(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagramEx() = %v, want %v", got, tt.want)
			}
			if got := isAnagramExtion(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagramExtion() = %v, want %v", got, tt.want)
			}
			if got := isAnagramExtionEx(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagramExtionEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 性能测试结果
// PS E:\gocode\trunk\src\go_exercise\leetcode\alg242> go test -bench=isAnagram -benchmem -cpu=1 -count=5
// goos: windows
// goarch: amd64
// pkg: go_exercise/leetcode/alg242
// Benchmark_isAnagram              4505256               274 ns/op              80 B/op          4 allocs/op
// Benchmark_isAnagram              4480234               261 ns/op              80 B/op          4 allocs/op
// Benchmark_isAnagram              4573382               256 ns/op              80 B/op          4 allocs/op
// Benchmark_isAnagram              4522174               255 ns/op              80 B/op          4 allocs/op
// Benchmark_isAnagram              4643635               257 ns/op              80 B/op          4 allocs/op
// Benchmark_isAnagramEx            3610980               320 ns/op             144 B/op          6 allocs/op
// Benchmark_isAnagramEx            3665695               318 ns/op             144 B/op          6 allocs/op
// Benchmark_isAnagramEx            3828115               329 ns/op             144 B/op          6 allocs/op
// Benchmark_isAnagramEx            3436594               326 ns/op             144 B/op          6 allocs/op
// Benchmark_isAnagramEx            3889620               328 ns/op             144 B/op          6 allocs/op
// Benchmark_isAnagramExtion        5352492               229 ns/op               0 B/op          0 allocs/op
// Benchmark_isAnagramExtion        5061422               225 ns/op               0 B/op          0 allocs/op
// Benchmark_isAnagramExtion        5147547               225 ns/op               0 B/op          0 allocs/op
// Benchmark_isAnagramExtion        5398894               223 ns/op               0 B/op          0 allocs/op
// Benchmark_isAnagramExtion        5157680               226 ns/op               0 B/op          0 allocs/op
// Benchmark_isAnagramExtionEx     83475356                14.8 ns/op             0 B/op          0 allocs/op
// Benchmark_isAnagramExtionEx     73314230                14.4 ns/op             0 B/op          0 allocs/op
// Benchmark_isAnagramExtionEx     80645161                14.3 ns/op             0 B/op          0 allocs/op
// Benchmark_isAnagramExtionEx     84918458                14.4 ns/op             0 B/op          0 allocs/op
// Benchmark_isAnagramExtionEx     80639740                14.4 ns/op             0 B/op          0 allocs/op
// PASS
// ok      go_exercise/leetcode/alg242     29.145s

func Benchmark_isAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isAnagram("anagram", "nagaram")
	}
}
func Benchmark_isAnagramEx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isAnagramEx("anagram", "nagaram")
	}
}
func Benchmark_isAnagramExtion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isAnagramExtion("anagram", "nagaram")
	}
}
func Benchmark_isAnagramExtionEx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isAnagramExtionEx("anagram", "nagaram")
	}
}
