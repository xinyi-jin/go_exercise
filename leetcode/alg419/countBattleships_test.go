package alg419

import "testing"

func Test_countBattleships(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_countBattleships_01", args{[][]byte{
			{'X', '.', '.', 'X'},
			{'.', '.', '.', 'X'},
			{'.', '.', '.', 'X'},
			{'.', '.', '.', '.'},
		},
		},
			2},
		{"Test_countBattleships_02", args{[][]byte{
			{'.'},
		},
		},
			0},
		{"Test_countBattleships_03", args{[][]byte{
			{'X', '.', '.', 'X'},
			{'.', '.', '.', 'X'},
			{'.', '.', '.', 'X'},
		},
		},
			2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBattleships(tt.args.board); got != tt.want {
				t.Errorf("%q. countBattleships() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
