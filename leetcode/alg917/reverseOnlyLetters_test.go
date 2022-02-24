package alg917

import "testing"

func Test_reverseOnlyLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_reverseOnlyLetters_01", args{"ab-cd"}, "dc-ba"},
		{"Test_reverseOnlyLetters_02", args{"a-bC-dEf-ghIj"}, "j-Ih-gfE-dCba"},
		{"Test_reverseOnlyLetters_03", args{"Test1ng-Leet=code-Q!"}, "Qedo1ct-eeLg=ntse-T!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseOnlyLetters(tt.args.s); got != tt.want {
				t.Errorf("reverseOnlyLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
