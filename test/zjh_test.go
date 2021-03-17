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
		// {Value: 12, Color: 0},
		// {Value: 11, Color: 0},
		// {Value: 12, Color: 1},
		{Value: 2, Color: 0},
		{Value: 3, Color: 0},
		{Value: 4, Color: 0},
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
