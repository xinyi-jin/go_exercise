package alg504

import "testing"

func Test_convertToBase7(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_convertToBase7_01", args{100}, "202"},
		{"Test_convertToBase7_02", args{-7}, "-10"},
		{"Test_convertToBase7_03", args{10}, "13"},
		{"Test_convertToBase7_04", args{-10}, "-13"},
		{"Test_convertToBase7_05", args{6}, "6"},
		{"Test_convertToBase7_06", args{0}, "0"},
		{"Test_convertToBase7_07", args{-1}, "-1"},
		{"Test_convertToBase7_08", args{-99999}, "-564354"},
		{"Test_convertToBase7_09", args{712479}, "6025125"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBase7(tt.args.num); got != tt.want {
				t.Errorf("convertToBase7() = %v, want %v", got, tt.want)
			}
			if got := convertToBase7Ex(tt.args.num); got != tt.want {
				t.Errorf("convertToBase7() = %v, want %v", got, tt.want)
			}
		})
	}
}
