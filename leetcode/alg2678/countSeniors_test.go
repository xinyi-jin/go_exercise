package alg2678

import "testing"

func Test_countSeniors(t *testing.T) {
	type args struct {
		details []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_countSeniors_01", args{[]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}}, 2},
		{"Test_countSeniors_02", args{[]string{"1313579440F2036", "2921522980M5644"}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSeniors(tt.args.details); got != tt.want {
				t.Errorf("countSeniors() = %v, want %v", got, tt.want)
			}
			if got := countSeniorsEx(tt.args.details); got != tt.want {
				t.Errorf("countSeniorsEx() = %v, want %v", got, tt.want)
			}
			if got := countSeniorsExtion(tt.args.details); got != tt.want {
				t.Errorf("countSeniorsExtion() = %v, want %v", got, tt.want)
			}
		})
	}
}
