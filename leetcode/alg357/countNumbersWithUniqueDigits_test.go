package alg357

import (
	"testing"
)

func Test_countNumbersWithUniqueDigits(t *testing.T) {
	type args struct {
		n int
	}
	var table = []int{1, 10, 91, 739, 5275, 32491, 168571, 712891, 2345851, 5611771, 8877691}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_countNumbersWithUniqueDigits_01", args{0}, table[0]},
		{"Test_countNumbersWithUniqueDigits_02", args{1}, table[1]},
		{"Test_countNumbersWithUniqueDigits_03", args{2}, table[2]},
		{"Test_countNumbersWithUniqueDigits_04", args{3}, table[3]},
		{"Test_countNumbersWithUniqueDigits_05", args{4}, table[4]},
		{"Test_countNumbersWithUniqueDigits_06", args{5}, table[5]},
		{"Test_countNumbersWithUniqueDigits_07", args{6}, table[6]},
		{"Test_countNumbersWithUniqueDigits_08", args{7}, table[7]},
		{"Test_countNumbersWithUniqueDigits_09", args{8}, table[8]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNumbersWithUniqueDigits(tt.args.n); got != tt.want {
				t.Errorf("countNumbersWithUniqueDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countNumbersWithUniqueDigitsEx(t *testing.T) {
	type args struct {
		n int
	}
	var table = []int{1, 10, 91, 739, 5275, 32491, 168571, 712891, 2345851, 5611771, 8877691}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_countNumbersWithUniqueDigitsEx_01", args{0}, table[0]},
		{"Test_countNumbersWithUniqueDigitsEx_02", args{1}, table[1]},
		{"Test_countNumbersWithUniqueDigitsEx_03", args{2}, table[2]},
		{"Test_countNumbersWithUniqueDigitsEx_04", args{3}, table[3]},
		{"Test_countNumbersWithUniqueDigitsEx_05", args{4}, table[4]},
		{"Test_countNumbersWithUniqueDigitsEx_06", args{5}, table[5]},
		{"Test_countNumbersWithUniqueDigitsEx_07", args{6}, table[6]},
		{"Test_countNumbersWithUniqueDigitsEx_08", args{7}, table[7]},
		{"Test_countNumbersWithUniqueDigitsEx_09", args{8}, table[8]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNumbersWithUniqueDigitsEx(tt.args.n); got != tt.want {
				t.Errorf("countNumbersWithUniqueDigitsEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
