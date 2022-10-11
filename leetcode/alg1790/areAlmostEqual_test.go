package alg1790

import "testing"

func Test_areAlmostEqual(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test_areAlmostEqual_01", args{"bank", "kanb"}, true},
		{"Test_areAlmostEqual_02", args{"attack", "defend"}, false},
		{"Test_areAlmostEqual_03", args{"kelb", "kelb"}, true},
		{"Test_areAlmostEqual_04", args{"abcd", "dcba"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areAlmostEqual(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("areAlmostEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
