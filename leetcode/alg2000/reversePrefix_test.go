package alg2000

import (
	"testing"
)

func Test_reversePrefix(t *testing.T) {
	type args struct {
		word string
		ch   byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_reversePrefix_01", args{"abcdefd", 'd'}, "dcbaefd"},
		{"Test_reversePrefix_02", args{"xyxzxe", 'z'}, "zxyxxe"},
		{"Test_reversePrefix_03", args{"abcd", 'z'}, "abcd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrefix(tt.args.word, tt.args.ch); got != tt.want {
				t.Errorf("reversePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reversePrefixEx(t *testing.T) {
	type args struct {
		word string
		ch   byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test_reversePrefixEx_01", args{"abcdefd", 'd'}, "dcbaefd"},
		{"Test_reversePrefixEx_02", args{"xyxzxe", 'z'}, "zxyxxe"},
		{"Test_reversePrefixEx_03", args{"abcd", 'z'}, "abcd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrefixEx(tt.args.word, tt.args.ch); got != tt.want {
				t.Errorf("reversePrefixEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
