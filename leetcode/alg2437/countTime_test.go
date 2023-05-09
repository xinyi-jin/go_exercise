package alg2437

import "testing"

func Test_countTime(t *testing.T) {
	type args struct {
		time string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_countTime_01", args{"?5:00"}, 2},
		{"Test_countTime_02", args{"0?:0?"}, 100},
		{"Test_countTime_03", args{"??:??"}, 1440},
		{"Test_countTime_04", args{"?2:16"}, 3},
		{"Test_countTime_05", args{"2?:??"}, 240},
		{"Test_countTime_06", args{"?9:?0"}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTime(tt.args.time); got != tt.want {
				t.Errorf("countTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
