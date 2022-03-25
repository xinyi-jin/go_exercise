package japansearchtable

import (
	"testing"
)

func Test_calcKey(t *testing.T) {
	type args struct {
		index []int16
		pos   []int16
	}
	var pos [14]int16
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"Test_calcKey_01", args{[]int16{
			1, 1, 1, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0,
		}, pos[0:]}, 0x0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcKey(tt.args.index, tt.args.pos); got != tt.want {
				t.Errorf("calcKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
