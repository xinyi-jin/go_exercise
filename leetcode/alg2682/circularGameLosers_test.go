package alg2682

import (
	"reflect"
	"testing"
)

func Test_circularGameLosers(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {"Test_circularGameLosers_01", args{5, 2}, []int{4, 5}},
		// {"Test_circularGameLosers_02", args{4, 4}, []int{2, 3, 4}},
		{"Test_circularGameLosers_03", args{6, 1}, []int{3, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := circularGameLosers(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("circularGameLosers() = %v, want %v", got, tt.want)
			}
		})
	}
}
