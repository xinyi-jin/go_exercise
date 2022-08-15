package alg640

import "testing"

func Test_solveEquation(t *testing.T) {
	type args struct {
		equation string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_solveEquation_01", args{"x+5-3+x=6+x-2"}, "x=2"},
		{"Test_solveEquation_02", args{"x=x"}, "Infinite solutions"},
		{"Test_solveEquation_03", args{"2x=x"}, "x=0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveEquation(tt.args.equation); got != tt.want {
				t.Errorf("solveEquation() = %v, want %v", got, tt.want)
			}
			if got := solveEquationEx(tt.args.equation); got != tt.want {
				t.Errorf("solveEquationEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
