package average

import (
	"testing"
)

func Test_averageOne(t *testing.T) {
	type args struct {
		a uint64
		b uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"Test_averageOne_01", args{0, 2}, 1},
		{"Test_averageOne_02", args{0, 1}, 0},
		{"Test_averageOne_03", args{1, 10}, 5},
		{"Test_averageOne_04", args{1, 11}, 6},
		{"Test_averageOne_05", args{18446744073709551615, 18446744073709551613}, 18446744073709551614},
		{"Test_averageOne_06", args{18446744073709551615, 18446744073709551614}, 18446744073709551614},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOne(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("averageOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_averageTwo(t *testing.T) {
	type args struct {
		a uint64
		b uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"Test_averageTwo_01", args{0, 2}, 1},
		{"Test_averageTwo_02", args{0, 1}, 0},
		{"Test_averageTwo_03", args{1, 10}, 5},
		{"Test_averageTwo_04", args{1, 11}, 6},
		{"Test_averageTwo_05", args{18446744073709551615, 18446744073709551613}, 18446744073709551614},
		{"Test_averageTwo_06", args{18446744073709551615, 18446744073709551614}, 18446744073709551614},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageTwo(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("averageTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_averageThree(t *testing.T) {
	type args struct {
		a uint64
		b uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"Test_averageThree_01", args{0, 2}, 1},
		{"Test_averageThree_02", args{0, 1}, 0},
		{"Test_averageThree_03", args{1, 10}, 5},
		{"Test_averageThree_04", args{1, 11}, 6},
		// {"Test_averageThree_05", args{18446744073709551615, 18446744073709551613}, 18446744073709551614},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageThree(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("averageThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
