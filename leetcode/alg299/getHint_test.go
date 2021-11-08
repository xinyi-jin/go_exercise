package alg299

import "testing"

func Test_getHint(t *testing.T) {
	type args struct {
		secret string
		guess  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test1", args{"1807", "7810"}, "1A3B"},
		{"Test2", args{"1123", "0111"}, "1A1B"},
		{"Test3", args{"1", "0"}, "0A0B"},
		{"Test4", args{"1", "1"}, "1A0B"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHint(tt.args.secret, tt.args.guess); got != tt.want {
				t.Errorf("%q. getHint() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
