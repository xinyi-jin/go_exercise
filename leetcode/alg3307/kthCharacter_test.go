package alg3307

import "testing"

func Test_kthCharacter(t *testing.T) {
	type args struct {
		k          int64
		operations []int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"Test_kthCharacter_01", args{5, []int{0, 0, 0}}, 'a'},
		{"Test_kthCharacter_02", args{10, []int{0, 1, 0, 1}}, byte('b')},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := kthCharacter(tt.args.k, tt.args.operations); got != tt.want {
			// 	t.Errorf("kthCharacter() = %v, want %v", got, tt.want)
			// }
			if got := kthCharacterEx(tt.args.k, tt.args.operations); got != tt.want {
				t.Errorf("kthCharacterEx() = %v, want %v", got, tt.want)
			}
		})
	}
}
