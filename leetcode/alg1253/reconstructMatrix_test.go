package alg1253

import (
	"reflect"
	"testing"
)

func Test_reconstructMatrix(t *testing.T) {
	type args struct {
		upper  int
		lower  int
		colsum []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Test_reconstructMatrix_01", args{2, 1, []int{1, 1, 1}}, [][]int{
			{1, 1, 0},
			{0, 0, 1},
		}},
		{"Test_reconstructMatrix_02", args{2, 3, []int{2, 2, 1, 1}}, [][]int{}},
		{"Test_reconstructMatrix_03", args{5, 5, []int{2, 1, 2, 0, 1, 0, 1, 2, 0, 1}}, [][]int{
			{1, 1, 1, 0, 1, 0, 0, 1, 0, 0},
			{1, 0, 1, 0, 0, 0, 1, 1, 0, 1},
		}},
		{"Test_reconstructMatrix_04", args{9, 2, []int{0, 1, 2, 0, 0, 0, 0, 0, 2, 1, 2, 1, 2}}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reconstructMatrix(tt.args.upper, tt.args.lower, tt.args.colsum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstructMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
