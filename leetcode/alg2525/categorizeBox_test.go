package alg2525

import "testing"

func Test_categorizeBox(t *testing.T) {
	type args struct {
		length int
		width  int
		height int
		mass   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_categorizeBox_01", args{1000, 35, 700, 300}, "Heavy"},
		{"Test_categorizeBox_02", args{200, 50, 800, 50}, "Neither"},
		{"Test_categorizeBox_03", args{2909, 3968, 3272, 727}, "Both"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := categorizeBox(tt.args.length, tt.args.width, tt.args.height, tt.args.mass); got != tt.want {
				t.Errorf("categorizeBox() = %v, want %v", got, tt.want)
			}
			if got := categorizeBoxEx(tt.args.length, tt.args.width, tt.args.height, tt.args.mass); got != tt.want {
				t.Errorf("categorizeBoxEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
