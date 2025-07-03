package alg3304

import "testing"

func Test_kthCharacter(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"Test_kthCharacter_01", args{5}, 'b'},
		{"Test_kthCharacter_02", args{10}, 'c'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthCharacter(tt.args.k); got != tt.want {
				t.Errorf("kthCharacter() = %v, want %v", got, tt.want)
			}
			if got := kthCharacterEx(tt.args.k); got != tt.want {
				t.Errorf("kthCharacterEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
