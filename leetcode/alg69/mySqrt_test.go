package alg69

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_mySqrt_01", args{4}, 2},
		{"Test_mySqrt_02", args{8}, 2},
		{"Test_mySqrt_03", args{0}, 0},
		{"Test_mySqrt_04", args{1}, 1},
		{"Test_mySqrt_05", args{2}, 1},
		{"Test_mySqrt_06", args{2147395600}, 46340},
		{"Test_mySqrt_07", args{2147395599}, 46339},
		{"Test_mySqrt_08", args{807810077}, 28421},
		{"Test_mySqrt_09", args{1464898982}, 38273},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
			if got := mySqrt2(tt.args.x); got != tt.want {
				t.Errorf("mySqrt2() = %v, want %v", got, tt.want)
			}
			if got := mySqrt3(tt.args.x); got != tt.want {
				t.Errorf("mySqrt3() = %v, want %v", got, tt.want)
			}
		})
	}
}
