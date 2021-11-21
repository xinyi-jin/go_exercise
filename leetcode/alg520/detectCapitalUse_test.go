package alg520

import "testing"

func Test_detectCapitalUse(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {"Test1", args{"USA"}, true},
		// {"Test2", args{"FlaG"}, false},
		{"Test3", args{"g"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCapitalUse(tt.args.word); got != tt.want {
				t.Errorf("%q. detectCapitalUse() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
