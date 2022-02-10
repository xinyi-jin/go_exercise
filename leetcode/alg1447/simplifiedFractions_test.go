package alg1447

import (
	"reflect"
	"testing"
)

func Test_simplifiedFractions(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Test_simplifiedFractions_01", args{2}, []string{"1/2"}},
		{"Test_simplifiedFractions_02", args{3}, []string{"1/2", "1/3", "2/3"}},
		{"Test_simplifiedFractions_03", args{4}, []string{"1/2", "1/3", "1/4", "2/3", "3/4"}},
		{"Test_simplifiedFractions_04", args{1}, []string{}},
		{"Test_simplifiedFractions_05", args{10}, []string{"1/2", "1/3", "1/4", "1/5", "1/6", "1/7", "1/8", "1/9", "1/10", "2/3", "2/5", "2/7", "2/9", "3/4", "3/5", "3/7", "3/8", "3/10", "4/5", "4/7", "4/9", "5/6", "5/7", "5/8", "5/9", "6/7", "7/8", "7/9", "7/10", "8/9", "9/10"}},
		{"Test_simplifiedFractions_06", args{15}, []string{"1/2", "1/3", "1/4", "1/5", "1/6", "1/7", "1/8", "1/9", "1/10", "1/11", "1/12", "1/13", "1/14", "1/15", "2/3", "2/5", "2/7", "2/9", "2/11", "2/13", "2/15", "3/4", "3/5", "3/7", "3/8", "3/10", "3/11", "3/13", "3/14", "4/5", "4/7", "4/9", "4/11", "4/13", "4/15", "5/6", "5/7", "5/8", "5/9", "5/11", "5/12", "5/13", "5/14", "6/7", "6/11", "6/13", "7/8", "7/9", "7/10", "7/11", "7/12", "7/13", "7/15", "8/9", "8/11", "8/13", "8/15", "9/10", "9/11", "9/13", "9/14", "10/11", "10/13", "11/12", "11/13", "11/14", "11/15", "12/13", "13/14", "13/15", "14/15"}},
		// {"Test_simplifiedFractions_07", args{100}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifiedFractions(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simplifiedFractions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSimplest(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test_isSimplest_01", args{5, 10}, false},
		{"Test_isSimplest_02", args{10, 15}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSimplest(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isSimplest() = %v, want %v", got, tt.want)
			}
		})
	}
}
