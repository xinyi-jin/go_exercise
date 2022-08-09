package alg1413

import "testing"

func Test_minStartValue(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_minStartValue_01", args{[]int{-3, 2, -3, 4, 2}}, 5},
		{"Test_minStartValue_02", args{[]int{1, 2}}, 1},
		{"Test_minStartValue_03", args{[]int{1, -2, -3}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minStartValue(tt.args.nums); got != tt.want {
				t.Errorf("minStartValue() = %v, want %v", got, tt.want)
			}
			if got := minStartValueEx(tt.args.nums); got != tt.want {
				t.Errorf("minStartValueEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
