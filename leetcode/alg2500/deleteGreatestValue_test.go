package alg2500

import "testing"

func Test_deleteGreatestValue(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_deleteGreatestValue_01", args{[][]int{
			{1, 2, 4},
			{3, 3, 1},
		}}, 8},
		{"Test_deleteGreatestValue_02", args{[][]int{{10}}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteGreatestValue(tt.args.grid); got != tt.want {
				t.Errorf("deleteGreatestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
