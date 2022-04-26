package chow

import (
	"hnmatch/gamerule/beard"
	"log"
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
		/* {"Test_searchChowCards_01", args{pool: [beard.BEARD_MAX]int{
			2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
			2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
		}, c: beard.BEARD_SMALL_7}, nil},
		{"Test_searchChowCards_02", args{pool: [beard.BEARD_MAX]int{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		}, c: beard.BEARD_SMALL_2}, nil},
		{"Test_searchChowCards_03", args{pool: [beard.BEARD_MAX]int{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		}, c: beard.BEARD_SMALL_7}, nil},
		{"Test_searchChowCards_04", args{pool: [beard.BEARD_MAX]int{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		}, c: beard.BEARD_SMALL_A}, nil},
		{"Test_searchChowCards_05", args{pool: [beard.BEARD_MAX]int{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		}, c: beard.BEARD_BIG_2}, nil},
		{"Test_searchChowCards_06", args{pool: [beard.BEARD_MAX]int{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		}, c: beard.BEARD_BIG_7}, nil},
		{"Test_searchChowCards_07", args{pool: [beard.BEARD_MAX]int{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		}, c: beard.BEARD_BIG_A}, nil}, */
		{"Test_searchChowCards_08", args{pool: [beard.BEARD_MAX]int{
			1, 3, 1, 0, 1, 1, 0, 2, 1, 1,
			2, 1, 1, 1, 1, 0, 1, 1, 1, 0,
		}, c: beard.BEARD_BIG_5}, nil},
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

func Test_27A(t *testing.T) {
	sourceCards := []int64{beard.BEARD_SMALL_2, beard.BEARD_SMALL_7, beard.BEARD_SMALL_A,
		beard.BEARD_BIG_2, beard.BEARD_BIG_7, beard.BEARD_BIG_A}
	for _, c := range sourceCards {
		cards := []int64{beard.BEARD_INVAILD, beard.BEARD_INVAILD, c}
		switch c % 10 {
		case 1:
			cards[0] = c + 5
			cards[1] = c + 8
		case 6:
			cards[0] = c - 5
			cards[1] = c + 3
		case 9:
			cards[0] = c - 8
			cards[1] = c - 3
		}
		log.Println(cards)
	}
}

func TestGetChowCards(t *testing.T) {
	type args struct {
		pool [beard.BEARD_MAX]int
		c    int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{"TestGetChowCards_01", args{pool: [beard.BEARD_MAX]int{
			1, 3, 1, 0, 1, 1, 0, 2, 1, 1,
			2, 1, 1, 1, 1, 0, 1, 1, 1, 0,
		}, c: beard.BEARD_BIG_5}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetChowCards(tt.args.pool, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChowCards() = %v, want %v", got, tt.want)
			}
		})
	}
}
