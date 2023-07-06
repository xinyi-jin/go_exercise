package alg2178

import (
	"reflect"
	"testing"
)

func Test_maximumEvenSplit(t *testing.T) {
	type args struct {
		finalSum int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{"Test_maximumEvenSplit_01", args{12}, []int64{}},
		{"Test_maximumEvenSplit_02", args{7}, []int64{}},
		{"Test_maximumEvenSplit_03", args{28}, []int64{}},
		{"Test_maximumEvenSplit_04", args{985612}, []int64{}},
		{"Test_maximumEvenSplit_05", args{2}, []int64{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := maximumEvenSplit(tt.args.finalSum); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("maximumEvenSplit() = %v, want %v", got, tt.want)
			// }
			got1 := maximumEvenSplit(tt.args.finalSum)
			got2 := maximumEvenSplitEx(tt.args.finalSum)
			got3 := maximumEvenSplitExtion(tt.args.finalSum)
			if !reflect.DeepEqual(got1, got2) || !reflect.DeepEqual(got1, got3) {
				t.Errorf("maximumEvenSplit() = %v, %v, %v", got1, got2, got3)
			}
		})
	}
}
func Benchmark_maximumEvenSplitEx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maximumEvenSplitEx(12)
	}
}
func Benchmark_maximumEvenSplitExtion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maximumEvenSplitExtion(12)
	}
}
