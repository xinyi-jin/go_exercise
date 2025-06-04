package alg3403

import "testing"

func Test_answerString(t *testing.T) {
	type args struct {
		word       string
		numFriends int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_answerString_01", args{"dbca", 2}, "dbc"},
		{"Test_answerString_02", args{"gggg", 4}, "g"},
		{"Test_answerString_03", args{"gggg", 1}, "gggg"},
		{"Test_answerString_04", args{"bif", 2}, "if"},
		{"Test_answerString_05", args{"aann", 2}, "nn"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := answerString(tt.args.word, tt.args.numFriends); got != tt.want {
				t.Errorf("answerString() = %v, want %v", got, tt.want)
			}
		})
	}
}
