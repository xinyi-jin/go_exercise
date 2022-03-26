package searchtable

import (
	"reflect"
	"testing"
)

func Test_getCardsValue(t *testing.T) {
	type args struct {
		cards []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test_getCardsValue_01", args{g_cards}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCardsValue(tt.args.cards); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCardsValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
