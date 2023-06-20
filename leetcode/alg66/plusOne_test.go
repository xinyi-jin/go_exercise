package alg66

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test_plusOne_01", args{[]int{1, 2, 3}}, []int{1, 2, 4}},
		{"Test_plusOne_02", args{[]int{1, 9, 9}}, []int{2, 0, 0}},
		{"Test_plusOne_03", args{[]int{9, 9, 9, 1, 0, 2, 9, 8, 9}}, []int{9, 9, 9, 1, 0, 2, 9, 9, 0}},
		{"Test_plusOne_04", args{[]int{9, 9, 9}}, []int{1, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
