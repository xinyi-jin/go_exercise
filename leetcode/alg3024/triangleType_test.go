package alg3024

import "testing"

func Test_triangleType(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_triangleType_01", args{[]int{3, 3, 3}}, "equilateral"},
		{"Test_triangleType_02", args{[]int{3, 4, 5}}, "scalene"},
		{"Test_triangleType_03", args{[]int{2, 4, 6}}, "none"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleType(tt.args.nums); got != tt.want {
				t.Errorf("triangleType() = %v, want %v", got, tt.want)
			}
		})
	}
}
