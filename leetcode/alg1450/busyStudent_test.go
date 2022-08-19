package alg1450

import "testing"

func Test_busyStudent(t *testing.T) {
	type args struct {
		startTime []int
		endTime   []int
		queryTime int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_busyStudent_01", args{[]int{1, 2, 3}, []int{3, 2, 7}, 4}, 1},
		{"Test_busyStudent_02", args{[]int{4}, []int{4}, 4}, 1},
		{"Test_busyStudent_03", args{[]int{4}, []int{4}, 5}, 0},
		{"Test_busyStudent_04", args{[]int{1, 1, 1, 1}, []int{1, 3, 2, 4}, 7}, 0},
		{"Test_busyStudent_05", args{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{10, 10, 10, 10, 10, 10, 10, 10, 10}, 5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := busyStudent(tt.args.startTime, tt.args.endTime, tt.args.queryTime); got != tt.want {
				t.Errorf("busyStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}
