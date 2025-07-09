package alg119

import (
	"reflect"
	"testing"
)

func Test_getRow(t *testing.T) {
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test_getRow_01", args{3}, []int{1, 3, 3, 1}},
		{"Test_getRow_02", args{0}, []int{1}},
		{"Test_getRow_03", args{1}, []int{1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRow(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
			if got := getRowEx(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRowEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeTable(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestMakeTable"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MakeTable()
		})
	}
}
