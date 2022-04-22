package chow

import (
	"hnmatch/gamerule/beard"
	"reflect"
	"testing"
)

func Test_searchChowCards(t *testing.T) {
	type args struct {
		pool [beard.BEARD_MAX]int
		c    int64
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"Test_searchChowCards_01", args{pool: [beard.BEARD_MAX]int{
			2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
			2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
		}, c: beard.BEARD_SMALL_7}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchChowCards(tt.args.pool, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchChowCards() = %v, want %v", got, tt.want)
				got.Traver()
			}
		})
	}
}
