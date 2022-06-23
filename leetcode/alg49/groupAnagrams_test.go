package alg49

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]string
	}{
		{"Test_01", args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}}, [][]string{
			{"bat"},
			{"nat", "tan"},
			{"ate", "eat", "tea"},
		}},
		{"Test_02", args{[]string{""}}, [][]string{{""}}},
		{"Test_02", args{[]string{"a"}}, [][]string{{"a"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := groupAnagrams(tt.args.strs); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("groupAnagrams() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
