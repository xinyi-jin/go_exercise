package chow

import (
	"testing"
)

func Test_inSequence(t *testing.T) {
	type args struct {
		card int64
		pool [34]int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test_inSequence_01", args{0, [34]int64{1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1},
		}, true},
		{"Test_inSequence_02", args{3, [34]int64{1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1},
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inSequence(tt.args.card, tt.args.pool); got != tt.want {
				t.Errorf("%q. inSequence() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
