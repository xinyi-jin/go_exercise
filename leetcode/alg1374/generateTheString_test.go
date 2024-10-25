package alg1374

import "testing"

func Test_generateTheString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_generateTheString_01", args{4}, "aaab"},
		{"Test_generateTheString_02", args{2}, "ab"},
		{"Test_generateTheString_03", args{7}, "aaaaaaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateTheString(tt.args.n); got != tt.want {
				t.Errorf("generateTheString() = %v, want %v", got, tt.want)
			}
			if got := generateTheStringEx(tt.args.n); got != tt.want {
				t.Errorf("generateTheStringEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
