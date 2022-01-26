package test

import (
	"fmt"
	"go_exercise/leetcode/zjh"
	"testing"
)

var pokerData []*zjh.Poker

func init() {
	pokerData = make([]*zjh.Poker, zjh.HandPokerNum)
	pokerData = []*zjh.Poker{
		{Value: 12, Color: 0},
		{Value: 11, Color: 0},
		{Value: 10, Color: 1},
	}
}

func TestZjh(t *testing.T) {
}

func TestCalcHandPokerType(T *testing.T) {

	hpt := zjh.CalcHandPokerType(pokerData, zjh.HandPokerNum)
	fmt.Printf("HandPokerType:%v \n%s", hpt, zjh.PokerDataInfo(pokerData))

}
func TestCalcHandPokerLogicValue(T *testing.T) {
	// A♦ K♦ A♣
	hplv := zjh.CalcHandPokerLogicValue(pokerData, zjh.HandPokerNum, zjh.POKERTYPE_DUIZI)
	fmt.Println("HandPokerLogicValue:", hplv)
}

func TestPokerDataSortShunZiA23(T *testing.T) {
	zjh.PokerDataSortShunZiA23(pokerData, zjh.HandPokerNum)
	fmt.Printf("PokerDataSortShunZiA23 %s", zjh.PokerDataInfo(pokerData))
}

func TestProb(t *testing.T) {
	// 高牌		对子	顺子	金花	顺金	豹子	组合总数
	// 16440	3744	720		1096	48		52		22100
	fmt.Printf("%0.4f %0.4f %0.4f %0.4f %0.4f %0.4f %0.4f",
		float32(10000*16440/22100.0000/10000),
		float32(10000*3744/22100.0000/10000),
		float32(10000*720/22100.0000/10000),
		float32(10000*1096/22100.0000/10000),
		float32(10000*48/22100.0000/10000),
		float32(10000*52/22100.0000/10000),

		float32(10000*16440/22100.0000/10000)+
			float32(10000*3744/22100.0000/10000)+
			float32(10000*720/22100.0000/10000)+
			float32(10000*1096/22100.0000/10000)+
			float32(10000*48/22100.0000/10000)+
			float32(10000*52/22100.0000/10000))
}

func TestCalcWinProb(t *testing.T) {
	zjh.CalcWinProb(zjh.PokerProbData, pokerData, zjh.HandPokerNum)
	fmt.Printf("%s", zjh.PokerDataInfo(pokerData))
}

func TestComparePoker(t *testing.T) {
	res := zjh.ComparePoker(
		[]*zjh.Poker{
			{Value: 9, Color: 0},
			{Value: 9, Color: 2},
			{Value: 10, Color: 1},
		},
		[]*zjh.Poker{
			{Value: 8, Color: 1},
			{Value: 9, Color: 3},
			{Value: 9, Color: 1},
		}, zjh.HandPokerNum)
	fmt.Println("compare poker ", res)
}
