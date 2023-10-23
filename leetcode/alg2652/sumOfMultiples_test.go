package alg2652

import "testing"

func Test_sumOfMultiples(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_sumOfMultiples_01", args{7}, 21},
		{"Test_sumOfMultiples_02", args{10}, 40},
		{"Test_sumOfMultiples_03", args{9}, 30},
		{"Test_sumOfMultiples_04", args{15}, 81},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfMultiples(tt.args.n); got != tt.want {
				t.Errorf("sumOfMultiples() = %v, want %v", got, tt.want)
			}
			if got := sumOfMultiplesEx(tt.args.n); got != tt.want {
				t.Errorf("sumOfMultiplesEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
