package ALG

import (
	"fmt"
	"go_exercise/tools"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

// 牌型		min		max		牌型

// 豹子		222		AAA		5
// 顺金		234		AKQ		4
// 金花		352		AKJ		3
// 顺子		234		AKQ		2
// 对子		223		AAK		1
// 高牌		235		AKJ		0

// 牌型
const (
	POKERTYPE_UNKNOW = -1
	POKERTYPE_GAOPAI = iota
	POKERTYPE_DUIZI
	POKERTYPE_SHUNZI
	POKERTYPE_JINHUA
	POKERTYPE_SHUNJIN
	POKERTYPE_BAOZI
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
	HandPokerNum  = 3  //玩家手牌数量
)

var (
	CNT_GAOPAI  int64
	CNT_DUIZI   int64
	CNT_SHUNZI  int64
	CNT_JINHUA  int64
	CNT_SHUNJIN int64
	CNT_BAOZI   int64
)

var PokerData []*Poker
var PokerLogicValueMap map[int64][]*Poker

type Poker struct {
	Value int64
	Color int64
}

// PokerDataInit 初始化牌库
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

// PokerDataShuffle 洗牌
func PokerDataShuffle(pokerData []*Poker) []*Poker {
	rand.Seed(time.Now().UnixNano())
	num := len(pokerData)
	for i := 0; i < num; i++ {
		n := rand.Intn(num - i)
		pokerData[i], pokerData[n] = pokerData[n], pokerData[i]
	}
	return pokerData
}

// PokerDataInfo 输出手牌信息
func PokerDataInfo(pokerData []*Poker) string {
	s := ""
	for _, v := range pokerData {
		if v != nil {
			s += fmt.Sprintf("%v ", PokerInfo(v))
		}
	}
	s += fmt.Sprintf("PokerDataInfo: %s \n", s)
	return s
}

// PokerInfo 输出单张牌信息
func PokerInfo(poker *Poker) string {
	p := ""
	switch poker.Value {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
		v := strconv.FormatInt(poker.Value+2, 10)
		p += v
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

// GetHandPokerTypeName 获取手牌类型名称
func GetHandPokerTypeName(hpt int) string {
	var s string
	switch hpt {
	case POKERTYPE_GAOPAI:
		s = "高牌"
	case POKERTYPE_DUIZI:
		s = "对子"
	case POKERTYPE_SHUNZI:
		s = "顺子"
	case POKERTYPE_JINHUA:
		s = "金花"
	case POKERTYPE_SHUNJIN:
		s = "顺金"
	case POKERTYPE_BAOZI:
		s = "豹子"
	default:
		s = "未知牌型"
	}
	return s
}

// HandPokerLogicValueInfo 输出牌力信息
func HandPokerLogicValueInfo(pokerLogicValueMap map[int64][]*Poker) {
	if len(pokerLogicValueMap) < 1 {
		return
	}
	logicValueArr := make(tools.Int64Slice, len(pokerLogicValueMap))
	pos := 0

	for k, _ := range pokerLogicValueMap {
		logicValueArr[pos] = k
		pos++
	}
	sort.Sort(logicValueArr)
	for _, v := range logicValueArr {
		fmt.Printf("HandPokerLogicValueInfo logicValue %d %s", v, PokerDataInfo(pokerLogicValueMap[v]))
	}
}

// AllPokerConbine 所有牌型组合
func AllPokerConbine(pokerData []*Poker, handPokerNum int) {
	// file, err := os.OpenFile("E:/zjh.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0)
	// defer file.Close()
	// if err != nil {
	// 	fmt.Println("打开文件失败")
	// 	return
	// }
	PokerLogicValueMap = make(map[int64][]*Poker)
	pokers := make([]*Poker, handPokerNum)
	pos := int64(0)
	value := len(PokerData)
	for i := 0; i < value; i++ {
		for j := i + 1; j < value; j++ {
			for z := j + 1; z < value; z++ {
				// pokers[0] = pokerData[i]
				// pokers[1] = pokerData[j]
				// pokers[2] = pokerData[z]
				// PokerLogicValueMap[pos] = pokers

				pokers = []*Poker{pokerData[i], pokerData[j], pokerData[z]}
				hpt := CalcHandPokerType(pokers, handPokerNum) // 计算牌型

				// 记录对应牌型数量
				WritePokerTypeCnt(hpt)

				// hplv := CalcHandPokerLogicValue(pokers, handPokerNum, hpt) // 计算牌力
				hplv := CalcPokerLogicValue(pokers, handPokerNum, hpt) // 计算牌力

				if _, ok := PokerLogicValueMap[hplv]; !ok {
					PokerLogicValueMap[hplv] = pokers
				} else {
					// fmt.Fprintf(file, "==== %d %d %s", hplv, hpt, PokerDataInfo(PokerLogicValueMap[hplv]))
				}
				pos += 1 // 统计牌型总数
			}
		}
	}
	// cnt := 0
	// for range PokerLogicValueMap {
	// 	cnt++
	// }
	HandPokerLogicValueInfo(PokerLogicValueMap)
	fmt.Println("AllPokerConbine pos", pos)
	fmt.Println("PokerLogicValueMap cnt ", len(PokerLogicValueMap)) // 牌力 牌型映射数量
}

// CalcHandPokerType 计算手牌牌型
func CalcHandPokerType(pokerData []*Poker, handPokerNum int) int {
	if len(pokerData) != handPokerNum {
		return POKERTYPE_UNKNOW
	}
	pokerDataTemp := make([]*Poker, handPokerNum)
	copy(pokerDataTemp, pokerData)

	// 豹子
	if IsBaoZi(pokerDataTemp, handPokerNum) {
		return POKERTYPE_BAOZI
	}
	// 顺金
	if IsShunJin(pokerDataTemp, handPokerNum) {
		return POKERTYPE_SHUNJIN
	}
	// 金花
	if IsJinHua(pokerDataTemp, handPokerNum) {
		return POKERTYPE_JINHUA
	}
	// 顺子
	if IsShunZi(pokerDataTemp, handPokerNum) {
		return POKERTYPE_SHUNZI
	}
	// 对子
	if IsDuiZi(pokerDataTemp, handPokerNum) {
		return POKERTYPE_DUIZI
	}
	// 高牌
	if IsGaoPai(pokerDataTemp, handPokerNum) {
		return POKERTYPE_GAOPAI
	}
	return POKERTYPE_UNKNOW
}

// CalcHandPokerLogicValue 计算手牌牌力值  9-12位 存储牌型 1-8位 存储牌值 牌值=牌型+花色
func CalcHandPokerLogicValue(pokerData []*Poker, handPokerNum, handPokerType int) int64 {
	var logicValue int64
	pokerDataTemp := make([]*Poker, handPokerNum)
	copy(pokerDataTemp, pokerData)

	PokerDataSort(pokerDataTemp)

	switch handPokerType {
	case POKERTYPE_DUIZI:
		PokerDataSortDuiZi(pokerDataTemp, handPokerNum)
		fallthrough
	case POKERTYPE_BAOZI, POKERTYPE_SHUNJIN, POKERTYPE_JINHUA, POKERTYPE_SHUNZI:
		logicValue = int64(handPokerType)<<11 + pokerDataTemp[0].Value<<7
	case POKERTYPE_GAOPAI:
		logicValue = int64(handPokerType)<<11 + pokerDataTemp[0].Value<<7 + pokerDataTemp[1].Value<<4 + pokerDataTemp[2].Value
	default:
		logicValue = -1
	}
	return logicValue
}

// CalcPokerLogicValue 计算牌力值  1-5 6-10 11-15 存储牌值和花色(2-5位存储牌值, 1存储花色--先比大小, 再比花色) 16-20 存储牌型
func CalcPokerLogicValue(pokerData []*Poker, handPokerNum, handPokerType int) int64 {
	var logicValue int64
	pokerDataTemp := make([]*Poker, handPokerNum)
	copy(pokerDataTemp, pokerData)

	PokerDataSort(pokerDataTemp)

	switch handPokerType {
	case POKERTYPE_DUIZI:
		PokerDataSortDuiZi(pokerDataTemp, handPokerNum)
		fallthrough
	case POKERTYPE_BAOZI, POKERTYPE_SHUNJIN, POKERTYPE_JINHUA, POKERTYPE_SHUNZI, POKERTYPE_GAOPAI:
		// 牌型 1牌值 1花色 2牌值 2花色 3牌值 3花色
		logicValue = int64(handPokerType)<<15 +
			pokerDataTemp[0].Value<<12 + pokerDataTemp[0].Color<<11 +
			pokerDataTemp[1].Value<<7 + pokerDataTemp[1].Color<<6 +
			pokerDataTemp[0].Value<<1 + pokerDataTemp[0].Color
	default:
		logicValue = -1
	}
	return logicValue
}

// IsBaoZi
func IsBaoZi(pokerData []*Poker, handPokerNum int) bool {
	if len(pokerData) != handPokerNum {
		return false
	}
	PokerDataSort(pokerData)
	v := pokerData[0].Value

	for i := 1; i < handPokerNum; i++ {
		if v != pokerData[i].Value {
			return false
		}
	}
	return true
}

// IsShunJin
func IsShunJin(pokerData []*Poker, handPokerNum int) bool {
	if len(pokerData) != handPokerNum {
		return false
	}
	if IsShunZi(pokerData, handPokerNum) && IsJinHua(pokerData, handPokerNum) {
		return true
	}
	return false
}

// IsJinHua
func IsJinHua(pokerData []*Poker, handPokerNum int) bool {
	if len(pokerData) != handPokerNum {
		return false
	}
	PokerDataSort(pokerData)
	c := pokerData[0].Color

	for i := 1; i < handPokerNum; i++ {
		if c != pokerData[i].Color {
			return false
		}
	}
	return true
}

// IsShunZi
func IsShunZi(pokerData []*Poker, handPokerNum int) bool {
	if len(pokerData) != handPokerNum {
		return false
	}
	PokerDataSort(pokerData)
	v := pokerData[0].Value

	// A23 也算作顺子
	if v == POKER_A && pokerData[1].Value == POKER_3 && pokerData[2].Value == POKER_2 {
		return true
	}

	for i := 1; i < handPokerNum; i++ {
		if v-1 != pokerData[i].Value {
			return false
		}
		v = pokerData[i].Value
	}
	return true
}

// IsDuiZi
func IsDuiZi(pokerData []*Poker, handPokerNum int) bool {
	if len(pokerData) != handPokerNum {
		return false
	}
	PokerDataSort(pokerData)

	for i := 0; i < handPokerNum-1; i++ {
		if pokerData[i].Value == pokerData[i+1].Value {
			return true
		}
	}
	return false
}

// IsGaoPai
func IsGaoPai(pokerData []*Poker, handPokerNum int) bool {
	if len(pokerData) != handPokerNum {
		return false
	}

	if IsBaoZi(pokerData, handPokerNum) || IsShunJin(pokerData, handPokerNum) || IsJinHua(pokerData, handPokerNum) ||
		IsShunZi(pokerData, handPokerNum) || IsDuiZi(pokerData, handPokerNum) {
		return false
	}
	return true
}

// PokerDataSort 按照牌值，花色 从大到小排序
func PokerDataSort(pokerData []*Poker) {
	l := len(pokerData)
	if l < 2 {
		return
	}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if pokerData[i].Value < pokerData[j].Value || (pokerData[i].Value == pokerData[j].Value && pokerData[i].Color < pokerData[j].Color) {
				pokerData[i], pokerData[j] = pokerData[j], pokerData[i]
			}
		}
	}
}

// PokerDataSortDuiZi 按照牌值，花色 从大到小排序--对子放最前边
func PokerDataSortDuiZi(pokerData []*Poker, handPokerNum int) {
	l := len(pokerData)
	if l < 2 {
		return
	}
	PokerDataSort(pokerData)

	for i := 0; i < handPokerNum-1; i++ {
		if pokerData[i].Value == pokerData[i+1].Value {
			if i != 0 {
				// 移动对子牌到开头
				pokerData[0], pokerData[i] = pokerData[i], pokerData[0]
				pokerData[1], pokerData[i+1] = pokerData[i+1], pokerData[1]
			}
			break
		}
	}
}

// 记录牌型数量
func WritePokerTypeCnt(hpt int) {
	switch hpt {
	case POKERTYPE_DUIZI:
		CNT_DUIZI++
	case POKERTYPE_BAOZI:
		CNT_BAOZI++
	case POKERTYPE_SHUNJIN:
		CNT_SHUNJIN++
	case POKERTYPE_JINHUA:
		CNT_JINHUA++
	case POKERTYPE_SHUNZI:
		CNT_SHUNZI++
	case POKERTYPE_GAOPAI:
		CNT_GAOPAI++
	default:
		fmt.Printf("WritePokerTypeCnt hpt %v", hpt)
	}
}

func init() {
	PokerDataInit(PokerValueNum, PokerColorNum, PokerNum)
	// fmt.Printf("PokerData init %s\n", PokerDataInfo(PokerData))
	if PokerData != nil {
		PokerDataShuffle(PokerData)
		// fmt.Printf("PokerData shuffle %s\n", PokerDataInfo(PokerData))
	}

	AllPokerConbine(PokerData, HandPokerNum)
	fmt.Printf("PokerTypeCnt %v %v %v %v %v %v",
		CNT_GAOPAI,
		CNT_DUIZI,
		CNT_SHUNZI,
		CNT_JINHUA,
		CNT_SHUNJIN,
		CNT_BAOZI)
}

// 炸金花算法实现
func zjh() {

}
