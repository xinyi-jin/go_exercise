package alg771

import "testing"

func Test_numJewelsInStones(t *testing.T) {
	type args struct {
		jewels string
		stones string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_numJewelsInStones_01", args{"aA", "aAAbbbb"}, 3},
		{"Test_numJewelsInStones_02", args{"z", "ZZ"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numJewelsInStones(tt.args.jewels, tt.args.stones); got != tt.want {
				t.Errorf("numJewelsInStones() = %v, want %v", got, tt.want)
			}
			if got := numJewelsInStonesEx(tt.args.jewels, tt.args.stones); got != tt.want {
				t.Errorf("numJewelsInStonesEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
