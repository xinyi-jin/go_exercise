package gdy

import (
	"fmt"
	"math/rand"
)

//牌序- K, Q, J,10, 9, 8, 7, 6, 5, 4, 3, 2, A
//     52  53
//黑桃-51,50,49,48,47,46,45,44,43,42,41,40,39
//红桃-38,37,36,35,34,33,32,31,30,29,28,27,26
//梅花-25,24,23,22,21,20,19,18,17,16,15,14,13
//方片-12,11,10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0
// 去掉癞子牌值
var PokerData = []int64{
	52, 53,
	51, 50, 49, 48, 47, 46, 45, 44, 43, 42, 41, 40, 39,
	38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26,
	25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13,
	12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0,
}
var PlayerNum = 4
var PokerCnt = int(54)

type GdyData struct {
	handCards    [][]int64
	playerNum    int
	bankerPos    int
	pokers       []int64
	surplusCnt   [16]int     // 0空闲 1-13对应牌值2-K 14对应小王 15对应大王
	idx          int         // 当前出牌位置
	outCards     []int64     // 已出牌数据
	expectResult [][]float64 // 预测概率
	otherPos     []int       // 下家 对家 上家
	bigCnt       [][]int64
	total        int64
}

func NewGdyData() *GdyData {
	return &GdyData{}
}

func (g *GdyData) init() {
	g.handCards = make([][]int64, PlayerNum)
	g.playerNum = PlayerNum
	g.bankerPos = rand.Intn(PlayerNum)
	g.pokers = make([]int64, len(PokerData))
	copy(g.pokers, PokerData)
	g.shuffle()
	g.idx = g.bankerPos
	g.dealOne(g.idx)
	g.surplusCntInit()
	g.surplusCntUpdateByHands(g.idx)
	g.outCards = make([]int64, 0)
	g.expectResult = make([][]float64, PlayerNum)
	g.otherPos = make([]int, PlayerNum-1)
	for i := 0; i < PlayerNum-1; i++ {
		g.otherPos[i] = (g.idx + i + 1) % PlayerNum
	}
	g.bigCnt = make([][]int64, PlayerNum)
	for i := 0; i < PlayerNum; i++ {
		g.bigCnt[i] = make([]int64, len(g.handCards[g.idx]))
	}
}

func (g *GdyData) surplusCntInit() {
	for _, c := range PokerData {
		if c >= 52 {
			if c == 52 {
				g.surplusCnt[14]++
			} else {
				g.surplusCnt[15]++
			}
		} else {
			if c%13 == 0 {
				g.surplusCnt[13]++
			} else {
				g.surplusCnt[c%13]++
			}
		}
	}
}

func (g *GdyData) surplusCntUpdateByHands(pos int) {
	cards := g.handCards[pos]
	for _, c := range cards {
		if c >= 52 {
			if c == 52 {
				g.surplusCnt[14]--
			} else {
				g.surplusCnt[15]--
			}
		} else {
			if c%13 == 0 {
				g.surplusCnt[13]--
			} else {
				g.surplusCnt[c%13]--
			}
		}
	}
}
func (g *GdyData) shuffle() {
	rand.Shuffle(len(g.pokers), func(i, j int) {
		g.pokers[i], g.pokers[j] = g.pokers[j], g.pokers[i]
	})
}

func (g *GdyData) dealOne(idx int) {
	for pos := 0; pos < PlayerNum; pos++ {
		if pos != idx {
			continue
		}
		if pos == g.bankerPos {
			g.handCards[pos] = g.pokers[len(g.pokers)-6:]
			g.pokers = g.pokers[:len(g.pokers)-6]
		} else {
			g.handCards[pos] = g.pokers[len(g.pokers)-5:]
			g.pokers = g.pokers[:len(g.pokers)-5]
		}
	}
}

// 生成测试结果
func (g *GdyData) genResult() {
	g.dealOther(g.idx)

	for _, pos := range g.otherPos {
		for i, c := range g.handCards[g.idx] {
			if c >= 52 || c%13 == 1 {
				continue
			}
			for _, v := range g.handCards[pos] {
				if v >= 52 || v%13 == 1 {
					continue
				}
				if c%13+1 == v%13 {
					g.bigCnt[pos][i]++
					break
				}
			}
		}
	}
	g.total++
}

func (g *GdyData) PrintResult() {
	// 打印预测结果
	/* for idx, pos := range g.otherPos {
		fmt.Printf("Expect idx:%v, pos:%v\n=======================\n", idx, pos)
		for i, c := range g.handCards[g.idx] {
			fmt.Printf("card:%v rate:%v%%\t", c, g.expectResult[pos][i]*100)
		}
		fmt.Println()
	}

	fmt.Println()

	// 打印测试结果
	for idx, pos := range g.otherPos {
		fmt.Printf("Test idx:%v, pos:%v\n=======================\n", idx, pos)
		for i, c := range g.handCards[g.idx] {
			fmt.Printf("card:%v rate:%v%%\t", c, float64(g.bigCnt[pos][i])/float64(g.total)*100)
		}
		fmt.Println()
	} */

	/* for idx, pos := range g.otherPos {
		fmt.Printf("idx:%v, pos:%v\n\n", idx, pos)
		for i, c := range g.handCards[g.idx] {
			fmt.Printf("Expect card:%v rate:%v%%\n", c, g.expectResult[pos][i]*100)
			fmt.Printf("Test   card:%v rate:%v%%\n\n", c, float64(g.bigCnt[pos][i])/float64(g.total)*100)
		}
	} */

	pos := g.otherPos[0]
	fmt.Printf("pos:%v\n\n", pos)
	maxDiff := float64(0)
	for i, c := range g.handCards[g.idx] {
		fmt.Printf("card:%v\n", c)
		fmt.Printf("Expect rate:%v%%\n", g.expectResult[pos][i]*100)
		fmt.Printf("Test   rate:%v%%\n", float64(g.bigCnt[pos][i])/float64(g.total)*100)
		diff := g.expectResult[pos][i]*100 - float64(g.bigCnt[pos][i])/float64(g.total)*100
		fmt.Printf("Diff   rate:%v%%\n\n", diff)
		if diff > maxDiff {
			maxDiff = diff
		}
	}
	fmt.Printf("MaxDiff   rate:%v%%\n\n", maxDiff)
}

// 预测结果
func (g *GdyData) expResult() {
	// 调用一下，便于之后使用每个人手牌数
	g.dealOther(g.idx)

	onePos := g.otherPos[0]
	g.expectResult[onePos] = make([]float64, len(g.handCards[g.idx]))
	for i, c := range g.handCards[g.idx] {
		g.expectResult[onePos][i] = g.expectSingle(onePos, c, []int{})
	}
	twoPos := g.otherPos[1]
	g.expectResult[twoPos] = make([]float64, len(g.handCards[g.idx]))
	for i, c := range g.handCards[g.idx] {
		g.expectResult[twoPos][i] = g.expectSingle(twoPos, c, []int{onePos})
	}
	threePos := g.otherPos[2]
	g.expectResult[threePos] = make([]float64, len(g.handCards[g.idx]))
	for i, c := range g.handCards[g.idx] {
		g.expectResult[threePos][i] = g.expectSingle(threePos, c, []int{onePos, twoPos})
	}

	// 打印预测结果
	/* for idx, pos := range g.otherPos {
		fmt.Printf("idx:%v, pos:%v\n=======================\n", idx, pos)
		for i, c := range g.handCards[g.idx] {
			fmt.Printf("card:%v rate:%v%%\t", c, g.expectResult[pos][i]*100)
		}
		fmt.Println()
	} */
}

func (g *GdyData) dealOther(idx int) {
	g.shuffle()
	temp := make([]int64, len(g.pokers))
	copy(temp, g.pokers)

	for _, pos := range g.otherPos {
		if pos == g.bankerPos {
			g.handCards[pos] = temp[len(temp)-6:]
			temp = temp[:len(temp)-6]
		} else {
			g.handCards[pos] = temp[len(temp)-5:]
			temp = temp[:len(temp)-5]
		}
	}
	/* for i := 1; i < PlayerNum; i++ {
		pos := (i + g.idx) % PlayerNum
		if pos == g.bankerPos {
			g.handCards = append(g.handCards, temp[len(temp)-6:])g.dealOther(g.idx)
			temp = temp[:len(temp)-6]
		} else {
			g.handCards = append(g.handCards, temp[len(temp)-5:])
			temp = temp[:len(temp)-5]
		}
	} */
}

func (g *GdyData) expectSingle(pos int, card int64, otherPos []int) float64 {
	res := float64(0)
	if card >= 52 || card%13 == 1 {
		return res
	}
	a := float64(len(g.handCards[pos]))
	b := float64(g.surplusCnt[card%13+1])
	c := float64(PokerCnt - len(g.outCards) - len(g.handCards[g.idx]))
	for _, pos := range otherPos {
		c -= float64(len(g.handCards[pos]))
	}
	return calc(a, b, c)
}

/*
  a.当前玩家手牌数
  b.剩余可压牌数
  c.剩余牌数
*/
func calc(a, b, c float64) float64 {
	return a * b / c
}
