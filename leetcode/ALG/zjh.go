package ALG

import (
	"fmt"
	"go_exercise/tools"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

const (
	// 高牌		对子	顺子	金花	顺金	豹子	组合总数
	// 16440	3744	720		1096	48		52		22100
	// 0.7439 	0.1694 	0.0326 	0.0496 	0.0022	0.0024	1.0000	概率

	POKERTYPE_PROB_GAOPAI  = 7439
	POKERTYPE_PROB_DUIZI   = 1694
	POKERTYPE_PROB_SHUNZI  = 326
	POKERTYPE_PROB_JINHUA  = 496
	POKERTYPE_PROB_SHUNJIN = 22
	POKERTYPE_PROB_BAOZI   = 24
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
	POKERTYPE_MAX
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
var PokerLogicValueMapSingle PokerLogicValueMap
var PokerProbData []PokerLogicValueMap
var PokerTypeProb = [POKERTYPE_MAX]int64{
	POKERTYPE_PROB_GAOPAI,
	POKERTYPE_PROB_DUIZI,
	POKERTYPE_PROB_SHUNZI,
	POKERTYPE_PROB_JINHUA,
	POKERTYPE_PROB_SHUNJIN,
	POKERTYPE_PROB_BAOZI,
}

type PokerLogicValueMap map[int64][]*Poker

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
	s := "PokerDataInfo:"
	for _, v := range pokerData {
		if v != nil {
			s += fmt.Sprintf("%v ", PokerInfo(v))
		}
	}
	s += fmt.Sprintf(" %s \n", s)
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

// PokerProbDataSort 单个牌型, 牌力排序
func PokerProbDataSort(PokerProbData []PokerLogicValueMap, hpt int, hplv int64) []int64 {
	logicValueArr := make(tools.Int64Slice, len(PokerProbData[hpt]))
	pos := 0
	for k, _ := range PokerProbData[hpt] {
		logicValueArr[pos] = k
		pos++
	}
	sort.Sort(logicValueArr)
	return logicValueArr
}

// AllPokerProbDataSort 所有牌型, 牌力排序 (废弃， 不必排, 可直接使用当前牌力值 / 牌力值总数 计算概率)
func AllPokerProbDataSort(plvm PokerLogicValueMap) []int64 {
	allLogicValueArr := make(tools.Int64Slice, len(plvm))
	pos := 0

	for k := range plvm {
		allLogicValueArr[pos] = k
		pos++
	}
	sort.Sort(allLogicValueArr)
	return allLogicValueArr
}

// AllPokerCombine 所有牌型组合
func AllPokerCombine(pokerData []*Poker, handPokerNum int) {
	// file, err := os.OpenFile("E:/zjh.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0)
	// defer file.Close()
	// if err != nil {
	// 	fmt.Println("打开文件失败")
	// 	return
	// }
	PokerLogicValueMapSingle = make(PokerLogicValueMap)
	pokers := make([]*Poker, handPokerNum)
	pos := int64(0)
	value := len(PokerData)
	for i := 0; i < value; i++ {
		for j := i + 1; j < value; j++ {
			for z := j + 1; z < value; z++ {
				// pokers[0] = pokerData[i]
				// pokers[1] = pokerData[j]
				// pokers[2] = pokerData[z]
				// PokerLogicValueMapSingle[pos] = pokers

				pokers = []*Poker{pokerData[i], pokerData[j], pokerData[z]}
				hpt := CalcHandPokerType(pokers, handPokerNum) // 计算牌型

				hplv := CalcHandPokerLogicValue(pokers, handPokerNum, hpt) // 计算牌力
				// hplv := CalcPokerLogicValue(pokers, handPokerNum, hpt) // 计算牌力

				if _, ok := PokerLogicValueMapSingle[hplv]; !ok {
					PokerLogicValueMapSingle[hplv] = pokers
				} /* else {
				fmt.Fprintf(file, "==== %d %d %s %s", hplv, hpt, PokerDataInfo(PokerLogicValueMapSingle[hplv]), PokerDataInfo(pokers))
				} */
				pos += 1 // 统计牌型总数

				// 记录对应牌型数量
				WritePokerTypeCnt(PokerProbData, pokers, hpt, hplv)
			}
		}
	}
	// cnt := 0
	// for range PokerLogicValueMapSingle {
	// 	cnt++
	// }
	// HandPokerLogicValueInfo(PokerLogicValueMapSingle)
	fmt.Println("AllPokerCombine pos", pos)
	fmt.Println("PokerLogicValueMapSingle cnt ", len(PokerLogicValueMapSingle)) // 牌力 牌型映射数量
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

// CalcHandPokerLogicValue 计算手牌牌力值  13-16位 存储牌型 1-12位 存储牌值 牌值=牌型+花色
func CalcHandPokerLogicValue(pokerData []*Poker, handPokerNum, handPokerType int) int64 {
	var logicValue int64
	pokerDataTemp := make([]*Poker, handPokerNum)
	copy(pokerDataTemp, pokerData)

	PokerDataSort(pokerDataTemp)

	switch handPokerType {
	case POKERTYPE_DUIZI:
		PokerDataSortDuiZi(pokerDataTemp, handPokerNum)
		logicValue = int64(handPokerType)<<12 + pokerDataTemp[0].Value<<7 + pokerData[2].Value
	case POKERTYPE_BAOZI, POKERTYPE_SHUNJIN, POKERTYPE_SHUNZI:
		// A32 重排
		if IsShunZiA23(pokerDataTemp, handPokerNum) {
			PokerDataSortShunZiA23(pokerDataTemp, handPokerNum)
		}
		logicValue = int64(handPokerType)<<12 + pokerDataTemp[0].Value<<7
	case POKERTYPE_GAOPAI, POKERTYPE_JINHUA:
		logicValue = int64(handPokerType)<<12 + pokerDataTemp[0].Value<<7 + pokerDataTemp[1].Value<<4 + pokerDataTemp[2].Value
	default:
		logicValue = -1
	}
	return logicValue
}

// CalcPokerLogicValue 计算牌力值  1-6 7-12 13-18 存储牌值和花色(3-6位存储牌值, 1-2存储花色--先比大小, 再比花色)   19-22 存储牌型
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
		// A32 重排
		if IsShunZiA23(pokerDataTemp, handPokerNum) {
			PokerDataSortShunZiA23(pokerDataTemp, handPokerNum)
		}
		// 牌型 1牌值 1花色 2牌值 2花色 3牌值 3花色
		logicValue = int64(handPokerType)<<18 +
			pokerDataTemp[0].Value<<14 + pokerDataTemp[0].Color<<12 +
			pokerDataTemp[1].Value<<8 + pokerDataTemp[1].Color<<6 +
			pokerDataTemp[2].Value<<2 + pokerDataTemp[2].Color
	default:
		logicValue = -1
	}
	return logicValue
}

// CalcWinProb 计算获胜概率  获胜概率 = 牌力百分比 牌型概率 和 牌值概率, 相同牌力算作输
func CalcWinProb(pokerProbData []PokerLogicValueMap, pokerData []*Poker, handPokerNum int) int64 {
	hpt := CalcHandPokerType(pokerData, handPokerNum)
	hplv := CalcHandPokerLogicValue(pokerData, handPokerNum, hpt)
	// fmt.Printf("CalcWinProb %d %d\n", hpt, hplv)

	hplvArr := PokerProbDataSort(pokerProbData, hpt, hplv) // 单个牌型牌力排序
	hplvCnt := len(hplvArr)
	index := 0
	for k, v := range hplvArr {
		if v == hplv {
			index = k
			break
		}
	}

	allhplvArr := AllPokerProbDataSort(PokerLogicValueMapSingle) //所有牌力排序
	allhplvCnt := len(allhplvArr)
	allIndex := 0
	for k, v := range allhplvArr {
		if v == hplv {
			allIndex = k
			break
		}
	}

	hptProb := PokerTypeProb[hpt]                // 牌型获胜概率
	hplvProb := index * 10000 / hplvCnt          // 同牌型下, 牌力获胜概率
	allhplvProb := allIndex * 10000 / allhplvCnt // 所有牌型下, 牌力获胜概率

	fmt.Printf("hptProb %v,hplvProb %v allhplvProb %v \n", hptProb, hplvProb, allhplvProb)
	return 0
}

// CalcWinCoinProb 计算赢分概率  赢分概率 = 赢牌概率 * 赢牌分 - 输牌概率 * 输牌分
func CalcWinCoinProb(pokerProbData []PokerLogicValueMap, pokerData []*Poker, handPokerNum int) int64 {

	return 0
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

// IsShunZiA23
func IsShunZiA23(pokerData []*Poker, handPokerNum int) bool {
	if len(pokerData) != handPokerNum {
		return false
	}
	PokerDataSort(pokerData)
	v := pokerData[0].Value

	// A23
	if v == POKER_A && pokerData[1].Value == POKER_3 && pokerData[2].Value == POKER_2 {
		return true
	}
	return false
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

// PokerDataSortShunZiA23 A23 排序
func PokerDataSortShunZiA23(pokerData []*Poker, handPokerNum int) {
	l := len(pokerData)
	if l < 2 {
		return
	}
	PokerDataSort(pokerData)

	var pokerATemp *Poker
	if pokerData[0].Value == POKER_A {
		pokerATemp = pokerData[0]
		pokerData = append(pokerData[:0], pokerData[1:]...)
	}
	pokerData = append(pokerData, pokerATemp)
}

// 记录牌型数量
func WritePokerTypeCnt(PokerProbData []PokerLogicValueMap, pokers []*Poker, hpt int, hplv int64) {
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
	// 记录牌型概率表
	PokerProbData[hpt][hplv] = pokers
}

func init() {
	PokerProbData = make([]PokerLogicValueMap, POKERTYPE_MAX)
	for i := 0; i < POKERTYPE_MAX; i++ {
		PokerProbData[i] = make(PokerLogicValueMap)
	}

	PokerDataInit(PokerValueNum, PokerColorNum, PokerNum)
	// fmt.Printf("PokerData init %s\n", PokerDataInfo(PokerData))
	if PokerData != nil {
		PokerDataShuffle(PokerData)
		// fmt.Printf("PokerData shuffle %s\n", PokerDataInfo(PokerData))
	}

	AllPokerCombine(PokerData, HandPokerNum)
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
