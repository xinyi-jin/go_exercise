package alg344

import (
	"fmt"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test_reverseString_01", args{[]byte{'h', 'e', 'l', 'l', 'o'}}},
		{"Test_reverseString_02", args{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printChar(tt.args.s)
			reverseString(tt.args.s)
			printChar(tt.args.s)
		})
	}
}
func printChar(c []byte) {
	s := ""
	for _, v := range c {
		s += fmt.Sprintf("%c, ", v)
	}
	fmt.Println(s)
}
