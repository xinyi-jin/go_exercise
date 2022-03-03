package alg258

import "testing"

func Test_addDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_addDigits_01", args{38}, 2},
		{"Test_addDigits_02", args{0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addDigits(tt.args.num); got != tt.want {
				t.Errorf("addDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addDigitsEx(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_addDigitsEx_01", args{38}, 2},
		{"Test_addDigitsEx_02", args{0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addDigitsEx(tt.args.num); got != tt.want {
				t.Errorf("addDigitsEx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addDigitsExtion(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test_addDigitsExtion_01", args{38}, 2},
		{"Test_addDigitsExtion_02", args{0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addDigitsExtion(tt.args.num); got != tt.want {
				t.Errorf("addDigitsExtion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAddDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addDigits(123812468)
	}
}
func BenchmarkAddDigitsEx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addDigitsEx(123812468)
	}
}
func BenchmarkAddDigitsExtion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addDigitsExtion(123812468)
	}
}
