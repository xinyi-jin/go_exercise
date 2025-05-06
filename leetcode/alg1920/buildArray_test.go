package alg1920

import (
	"reflect"
	"testing"
)

func Test_buildArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test_buildArray_01", args{[]int{0, 2, 1, 5, 3, 4}}, []int{0, 1, 2, 4, 5, 3}},
		{"Test_buildArray_02", args{[]int{5, 0, 1, 2, 3, 4}}, []int{4, 5, 0, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArray() = %v, want %v", got, tt.want)
			}
			if got := buildArrayEx(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArrayEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
