package rldqn

import "testing"

func TestEvaluation(t *testing.T) {
	mahjongData := NewMahjongData()
	mahjongData.Hands = append(mahjongData.Hands, []int64{
		MJ_BAMBOO_1,
		MJ_BAMBOO_2,
		MJ_BAMBOO_3,
		MJ_CHARACTERS_9,
		MJ_CHARACTERS_9,
	}...)
	mahjongData.Discards = append(mahjongData.Discards, []int64{
		MJ_CHARACTERS_5,
		MJ_CHARACTERS_5,
		MJ_CHARACTERS_5,
	}...)
	mahjongData.Piles[0] = append(mahjongData.Piles[0], []int64{
		MJ_BAMBOO_5,
		MJ_BAMBOO_5,
		MJ_BAMBOO_5,
	}...)
	type args struct {
		mahjongData *MahjongData
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestEvaluation_01", args{
			mahjongData: mahjongData,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Evaluation(tt.args.mahjongData)
		})
	}
}
