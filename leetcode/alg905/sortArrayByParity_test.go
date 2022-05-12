package alg905

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParity(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test_sortArrayByParity_01", args{[]int{3, 1, 2, 4}}, []int{}},
		{"Test_sortArrayByParity_02", args{[]int{0}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParity(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortArrayByParityEx(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test_sortArrayByParityEx_01", args{[]int{3, 1, 2, 4}}, []int{}},
		{"Test_sortArrayByParityEx_02", args{[]int{0}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParityEx(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParityEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
