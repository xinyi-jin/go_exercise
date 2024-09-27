package alg2516

import "testing"

func Test_takeCharacters(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_takeCharacters_01", args{"aabaaaacaabc", 2}, 8},
		{"Test_takeCharacters_02", args{"a", 0}, 0},
		{"Test_takeCharacters_03", args{"bcca", 1}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := takeCharacters(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("takeCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
