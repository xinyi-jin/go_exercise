package test

import (
	"fmt"
	"go_exercise/leetcode/ALG"
	_ "go_exercise/leetcode/ALG"
	"testing"
)

var pokerData []*ALG.Poker

func init() {
	pokerData = make([]*ALG.Poker, ALG.HandPokerNum)
	pokerData = []*ALG.Poker{
		{Value: 12, Color: 0},
		{Value: 11, Color: 0},
		{Value: 10, Color: 1},
	}
}

func TestZjh(t *testing.T) {
}

func TestCalcHandPokerType(T *testing.T) {

	hpt := ALG.CalcHandPokerType(pokerData, ALG.HandPokerNum)
	fmt.Printf("HandPokerType:%v \n%s", hpt, ALG.PokerDataInfo(pokerData))

}
func TestCalcHandPokerLogicValue(T *testing.T) {
	// A♦ K♦ A♣
	hplv := ALG.CalcHandPokerLogicValue(pokerData, ALG.HandPokerNum, ALG.POKERTYPE_DUIZI)
	fmt.Println("HandPokerLogicValue:", hplv)
}

func TestPokerDataSortShunZiA23(T *testing.T) {
	ALG.PokerDataSortShunZiA23(pokerData, ALG.HandPokerNum)
	fmt.Printf("PokerDataSortShunZiA23 %s", ALG.PokerDataInfo(pokerData))
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
	ALG.CalcWinProb(ALG.PokerProbData, pokerData, ALG.HandPokerNum)
	fmt.Printf("%s", ALG.PokerDataInfo(pokerData))
}
