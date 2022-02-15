package alg1380

import (
	"reflect"
	"testing"
)

func Test_luckyNumbers(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"Test_luckyNumbers_01", args{[][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}}, []int{15}},
		{"Test_luckyNumbers_02", args{[][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}}, []int{12}},
		{"Test_luckyNumbers_03", args{[][]int{{7, 8}, {1, 2}}}, []int{7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := luckyNumbers(tt.args.matrix); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("luckyNumbers() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
