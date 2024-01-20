package alg2788

import (
	"reflect"
	"testing"
)

func Test_splitWordsBySeparator(t *testing.T) {
	type args struct {
		words     []string
		separator byte
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Test_splitWordsBySeparator_01", args{[]string{"one.two.three", "four.five", "six"}, '.'}, []string{"one", "two", "three", "four", "five", "six"}},
		{"Test_splitWordsBySeparator_02", args{[]string{"$easy$", "$problem$"}, '$'}, []string{"easy", "problem"}},
		{"Test_splitWordsBySeparator_03", args{[]string{"|||"}, '|'}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitWordsBySeparator(tt.args.words, tt.args.separator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitWordsBySeparator() = %v, want %v", got, tt.want)
			}
			if got := splitWordsBySeparatorEx(tt.args.words, tt.args.separator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitWordsBySeparatorEx() = %v, want %v", got, tt.want)
			}
			if got := splitWordsBySeparatorExtion(tt.args.words, tt.args.separator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitWordsBySeparatorExtion() = %v, want %v", got, tt.want)
			}
		})
	}
}
