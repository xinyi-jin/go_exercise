package alg537

import "testing"

func Test_complexNumberMultiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_complexNumberMultiply_01", args{"1+1i", "1+1i"}, "0+2i"},
		{"Test_complexNumberMultiply_02", args{"1+-1i", "1+-1i"}, "0+-2i"},
		{"Test_complexNumberMultiply_03", args{"8+2i", "5+9i"}, "22+82i"},
		{"Test_complexNumberMultiply_04", args{"6+-2i", "5+9i"}, "48+44i"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := complexNumberMultiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("complexNumberMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
