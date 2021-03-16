package ALG

import (
	"fmt"
	"math/rand"
	"time"
)

// 牌值
const (
	POKER_2 = iota
	POKER_3
	POKER_4
	POKER_5
	POKER_6
	POKER_7
	POKER_8
	POKER_9
	POKER_10
	POKER_J
	POKER_Q
	POKER_K
	POKER_A
)

// 花色
const (
	COLOR_Diamond = iota
	COLOR_Club
	COLOR_Heart
	COLOR_Spade
)

var (
	PokerValueNum = 13
	PokerColorNum = 4
	PokerNum      = 52 //去除大小王的牌数
)

var PokerData []*Poker

type Poker struct {
	Value int64
	Color int64
}

func PokerDataInit(value, color, pokerNum int) {
	PokerData = make([]*Poker, pokerNum)
	pos := 0
	for i := 0; i < value; i++ {
		for j := 0; j < color; j++ {
			PokerData[pos] = &Poker{
				Value: int64(i),
				Color: int64(j),
			}
			pos++
		}
	}
}

func PokerDataShuffle(pokerData []*Poker) []*Poker {
	rand.Seed(time.Now().UnixNano())
	num := len(pokerData)
	for i := 0; i < num; i++ {
		n := rand.Intn(num - i)
		pokerData[i], pokerData[n] = pokerData[n], pokerData[i]
	}
	return pokerData
}

func PokerDataInfo(pokerData []*Poker) {
	fmt.Printf("PokerDataInfo:")
	for _, v := range pokerData {
		if v != nil {
			fmt.Printf("%v, ", PokerInfo(v))
		}
	}
}

func PokerInfo(poker *Poker) string {
	p := ""
	switch poker.Value {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
		p += string(rune(poker.Value + 2))
	case 9:
		p += "J"
	case 10:
		p += "Q"
	case 11:
		p += "K"
	case 12:
		p += "A"
	default:
		p += "novalue"
	}

	switch poker.Color {
	case 0:
		p += "♦"
	case 1:
		p += "♣"
	case 2:
		p += "♥"
	case 3:
		p += "♠"
	default:
		p += "nocard"
	}

	return p
}

func init() {
	PokerDataInit(PokerValueNum, PokerColorNum, PokerNum)
	fmt.Printf("PokerData init \n")
	PokerDataInfo(PokerData)
	if PokerData != nil {
		fmt.Printf("PokerData shuffle \n")
		PokerDataShuffle(PokerData)
		PokerDataInfo(PokerData)
	}
}

// 炸金花算法实现
func zjh() {

}
