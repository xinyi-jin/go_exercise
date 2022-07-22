package alg757

import "testing"

func Test_intersectionSizeTwo(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{"Test_intersectionSizeTwo_01", args{[][]int{
			{1, 3},
			{1, 4},
			{2, 5},
			{3, 5},
		}}, 3},
		{"Test_intersectionSizeTwo_02", args{[][]int{
			{1, 2},
			{2, 3},
			{2, 4},
			{4, 5},
		}}, 5},
		{"Test_intersectionSizeTwo_03", args{[][]int{
			{1, 20},
			{2, 3},
			{2, 4},
			{4, 5},
		}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := intersectionSizeTwo(tt.args.intervals); gotAns != tt.wantAns {
				t.Errorf("intersectionSizeTwo() = %v, want %v", gotAns, tt.wantAns)
			}
			if gotAns := intersectionSizeTwoEx(tt.args.intervals); gotAns != tt.wantAns {
				t.Errorf("intersectionSizeTwoEx() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
