package alg1353

import "testing"

func Test_maxEvents(t *testing.T) {
	type args struct {
		events [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_maxEvents_01", args{[][]int{{1, 2}, {2, 3}, {3, 4}}}, 3},
		{"Test_maxEvents_02", args{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}}}, 4},
		{"Test_maxEvents_03", args{[][]int{{1, 3}, {2, 3}, {3, 4}, {1, 2}}}, 4},
		{"Test_maxEvents_04", args{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}}, 4},
		{"Test_maxEvents_05", args{[][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}}}, 7},
		{"Test_maxEvents_06", args{[][]int{{1, 2}, {1, 2}, {3, 3}, {1, 5}, {1, 5}}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEvents(tt.args.events); got != tt.want {
				t.Errorf("maxEvents() = %v, want %v", got, tt.want)
			}
			if got := maxEventsEx(tt.args.events); got != tt.want {
				t.Errorf("maxEventsEx() = %v, want %v", got, tt.want)
			}
		})
	}
	// arr := []int{1, 2, 3, 4, 5}
	// arr = append(arr[:0], arr[1:]...)
	// t.Log(arr)
}
