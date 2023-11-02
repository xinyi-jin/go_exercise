package alg2103

import "testing"

func Test_countPoints(t *testing.T) {
	type args struct {
		rings string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_countPoints_01", args{"B0B6G0R6R0R6G9"}, 1},
		{"Test_countPoints_02", args{"B0R0G0R9R0B0G0"}, 1},
		{"Test_countPoints_03", args{"G4"}, 0},
		{"Test_countPoints_04", args{"G7G8G0"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPoints(tt.args.rings); got != tt.want {
				t.Errorf("countPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
