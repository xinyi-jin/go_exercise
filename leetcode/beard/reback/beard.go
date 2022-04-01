package reback

import (
	"log"
)

type HuHandResult struct {
	Sequence [7]*Group
	SeqCnt   int
}
type Group struct {
	Cards [4]int64 // 当前组牌值
	Flag  int64    // 组牌类型标记
	HuXi  int      // 胡息
}

type Item struct {
	card int64 // 牌值
	i    int   // 第几种拆法
	j    int   // 正在进行的第几种拆法
	huxi int   // 胡息数
}

func (hhr *HuHandResult) init() {
	for i := 0; i < len(hhr.Sequence); i++ {
		hhr.Sequence[i] = &Group{
			Cards: [4]int64{-1, -1, -1, -1},
			Flag:  -1,
			HuXi:  0,
		}
	}
	hhr.SeqCnt = -1
}

func (hhr *HuHandResult) push(cards [4]int64, flag int64) {
	if hhr.SeqCnt+1 == len(hhr.Sequence) {
		return
	}
	hhr.SeqCnt++
	hhr.Sequence[hhr.SeqCnt] = &Group{
		Cards: cards,
		Flag:  flag,
		HuXi:  getGroupHuXiByFlag(flag),
	}
}

func (hhr *HuHandResult) pop() {
	if hhr.SeqCnt == -1 {
		return
	}
	hhr.Sequence[hhr.SeqCnt] = &Group{
		Cards: [4]int64{-1, -1, -1, -1},
		Flag:  -1,
		HuXi:  0,
	}
	hhr.SeqCnt--
}

func getGroupHuXiByFlag(flag int64) int {
	// 将牌，绞牌没有胡息
	if flag == GROUPTYPE_JIANG || flag == GROUPTYPE_JIAO_aaA || flag == GROUPTYPE_JIAO_aAA {
		return 0
	}
	return BeardFlagHuXiMap[flag]
}

func getKanHuxi(pool []int) (int, *HuHandResult) {
	huxi := 0
	res := &HuHandResult{}
	res.init()

	for i := BEARD_SMALL_1; i < BEARD_MAX; i++ {
		if pool[i] == 3 {
			pool[i] = 0
			if i < BEARD_BIG_1 {
				huxi += HUXI_KAN_SMALL
				res.push([4]int64{i, i, i, -1}, GROUPTYPE_KAN_SMALL)
			} else {
				huxi += HUXI_KAN_BIG
				res.push([4]int64{i, i, i, -1}, GROUPTYPE_KAN_BIG)
			}
		}
	}
	return huxi, res
}

func getShunHuXiBigNot27A(pool []int, num int) (int, *HuHandResult) {
	tempPool := make([]int, 10)
	copy(tempPool, pool)
	tempPool[1] -= num
	tempPool[6] -= num
	tempPool[9] -= num

	otherShunNum := 0
	res := &HuHandResult{}
	res.init()

	for i := 0; i < 10; i++ {
		n := tempPool[i]
		if n == 0 {
			continue
		}
		if i+2 >= 10 {
			return -1, nil
		}
		if tempPool[i+1] < n || tempPool[i+2] < tempPool[i+1] {
			return -1, nil
		}
		tempPool[i+1] -= n
		tempPool[i+2] -= n

		if i > 0 {
			otherShunNum += n
			curCard := int64(i + 10)

			for i := 0; i < n; i++ {
				res.push([4]int64{curCard, curCard + 1, curCard + 2, -1}, GROUPTYPE_SHUN_BIG)
			}
		}
	}
	return num*HUXI_123_27A_BIG + otherShunNum*HUXI_SHUN_BIG, res
}

func getShunHuXiBig(pool []int) (int, *HuHandResult) {
	if pool[BEARD_BIG_1] > pool[BEARD_BIG_2] || pool[BEARD_BIG_1] > pool[BEARD_BIG_3] {
		return -1, nil
	}
	var res = &HuHandResult{}
	res.init()

	sum := 0
	for _, v := range pool[BEARD_BIG_1:] {
		sum += v
	}
	if sum == 0 {
		return 0, res
	}

	// 仅拆27A 和 顺子
	maxHuxi := -1
	tempPool := make([]int, 10)
	copy(tempPool, pool[BEARD_BIG_1:])
	n_123 := tempPool[0]
	tempPool[0] -= n_123
	tempPool[1] -= n_123
	tempPool[2] -= n_123

	for i := 0; i < 5; i++ {
		copy(tempPool, pool[BEARD_BIG_1:])
		tempPool[0] -= n_123
		tempPool[1] -= n_123
		tempPool[2] -= n_123
		if tempPool[1] < i || tempPool[6] < i || tempPool[9] < i {
			break
		}
		huXi27A := -1
		huXi123 := n_123 * HUXI_123_27A_BIG
		huXi27A, res = getShunHuXiBigNot27A(tempPool, i)
		if huXi27A >= 0 {
			if huXi123+huXi27A > maxHuxi {
				maxHuxi = huXi123 + huXi27A

				for n := 0; n < n_123; n++ {
					res.push([4]int64{BEARD_BIG_1, BEARD_BIG_2, BEARD_BIG_3, -1}, GROUPTYPE_123_27A_BIG)
				}
				for n := 0; n < i; n++ {
					res.push([4]int64{BEARD_BIG_2, BEARD_BIG_7, BEARD_BIG_A, -1}, GROUPTYPE_123_27A_BIG)
				}
			}
		}
	}
	if maxHuxi == -1 {
		res = nil
	}
	return maxHuxi, res
}

func getShunHuXiSmall(pool []int) (int, *HuHandResult) {
	items := [7]Item{}
	pos := -1 // 牌值索引
	find := true
	maxHuXi := -1
	curCard := BEARD_SMALL_1

	var res = &HuHandResult{}

	var huHandResult = &HuHandResult{}
	huHandResult.init()

	for {
		if find {
			find = false
			finded := false
			if pos == 6 {
				// 异常处理
				log.Fatalln("异常！！！")
			}

			for i := curCard; i < BEARD_BIG_1; i++ {
				if pool[i] > 0 {
					curCard = i
					pos++
					items[pos].card = curCard
					finded = true
					break
				}
			}

			if !finded {
				huXiSmall := 0
				for i := 0; i <= pos; i++ {
					huXiSmall += items[i].huxi
				}
				huXiBig, bigCombine := getShunHuXiBig(pool)
				if huXiBig >= 0 {
					if huXiSmall+huXiBig > maxHuXi {
						maxHuXi = huXiSmall + huXiBig

						res.init()
						for i := 0; i <= huHandResult.SeqCnt; i++ {
							res.push(huHandResult.Sequence[i].Cards, huHandResult.Sequence[i].Flag)
						}
						for i := 0; i <= bigCombine.SeqCnt; i++ {
							res.push(bigCombine.Sequence[i].Cards, bigCombine.Sequence[i].Flag)
						}
					}
				}
				goto reback
			}
		}

		switch items[pos].i {
		case 0: // 顺子
			items[pos].i = 1
			if curCard < 8 && pool[curCard+1] > 0 && pool[curCard+2] > 0 {
				items[pos].j = 1
				items[pos].card = curCard
				pool[curCard]--
				pool[curCard+1]--
				pool[curCard+2]--
				find = true
				flag := int64(-1)
				if curCard == BEARD_SMALL_1 {
					items[pos].huxi = HUXI_123_27A_SMALL
					flag = GROUPTYPE_123_27A_SMALL
				} else {
					items[pos].huxi = HUXI_SHUN_SMALL
					flag = GROUPTYPE_SHUN_SMALL
				}
				huHandResult.push([4]int64{curCard, curCard + 1, curCard + 2, -1}, flag)
				continue
			}
		case 1: // 小小大绞
			items[pos].i = 2
			if pool[curCard] == 2 && pool[curCard+10] > 0 {
				items[pos].j = 2
				pool[curCard] -= 2
				pool[curCard+10]--
				find = true
				huHandResult.push([4]int64{curCard, curCard, curCard + 10, -1}, GROUPTYPE_JIAO_aaA)
				continue
			}
		case 2: // 大大小绞
			items[pos].i = 3
			if pool[curCard+10] == 2 {
				items[pos].j = 3
				pool[curCard+10] -= 2
				pool[curCard]--
				find = true
				huHandResult.push([4]int64{curCard, curCard + 10, curCard + 10, -1}, GROUPTYPE_JIAO_aAA)
				continue
			}
		case 3:
			items[pos].i = 4
			if curCard == BEARD_SMALL_2 && pool[BEARD_SMALL_7] > 0 && pool[BEARD_SMALL_A] > 0 {
				items[pos].j = 4
				items[pos].huxi = HUXI_123_27A_SMALL
				pool[BEARD_SMALL_2]--
				pool[BEARD_SMALL_7]--
				pool[BEARD_SMALL_A]--
				find = true
				huHandResult.push([4]int64{BEARD_SMALL_2, BEARD_SMALL_7, BEARD_SMALL_A, -1}, GROUPTYPE_123_27A_SMALL)
				continue
			}
		}

	reback: // 回溯
		if pos < 0 {
			goto finish
		}

		// 没找到 或者 已经是最后一个组合了
		if items[pos].i == 0 || items[pos].i == 4 && items[pos].j != 4 {
			items[pos] = Item{}
			if pos == 0 {
				goto finish
			}
			pos--
			huHandResult.pop()
			goto reback
		}

		curCard = items[pos].card
		if items[pos].i > 0 && items[pos].i == items[pos].j {
			items[pos].huxi = 0
			switch items[pos].j {
			case 1:
				pool[curCard]++
				pool[curCard+1]++
				pool[curCard+2]++
			case 2:
				pool[curCard] += 2
				pool[curCard+10]++
			case 3:
				pool[curCard+10] += 2
				pool[curCard]++
			case 4:
				pool[1]++
				pool[6]++
				pool[9]++
				items[pos] = Item{}
				pos--
				huHandResult.pop()
				if pos == 0 {
					goto finish
				} else {
					goto reback
				}
			}
		}
	}

finish:
	if maxHuXi == -1 {
		res = nil
	}
	return maxHuXi, res
}

// 计算胡息
func CalcHuXi(pool []int, handsNum int) (int, *HuHandResult) {
	var res = &HuHandResult{}
	res.init()

	huXiKan, kanCombine := getKanHuxi(pool) // 坎牌胡息数

	// 不带将牌
	if handsNum%3 == 0 {
		otherHuXi := -1
		otherHuXi, res = getShunHuXiSmall(pool)
		if otherHuXi < 0 {
			return -1, nil
		}
		for i := 0; i <= kanCombine.SeqCnt; i++ {
			res.push(kanCombine.Sequence[i].Cards, kanCombine.Sequence[i].Flag)
		}
		return huXiKan + otherHuXi, res
	}

	// 带将牌
	max := -1
	tempPool := make([]int, BEARD_MAX)
	for i := int64(0); i < BEARD_MAX; i++ {
		if pool[i] > 1 {
			copy(tempPool, pool)
			tempPool[i] -= 2

			otherHuXi := -1
			res = &HuHandResult{}
			res.init()
			otherHuXi, res = getShunHuXiSmall(tempPool)

			if otherHuXi < 0 {
				continue
			}
			if otherHuXi > max {
				res.push([4]int64{i, i, -1, -1}, GROUPTYPE_JIANG)
				max = otherHuXi
			}
		}
	}
	if max < 0 {
		return -1, nil
	}

	for i := 0; i <= kanCombine.SeqCnt; i++ {
		res.push(kanCombine.Sequence[i].Cards, kanCombine.Sequence[i].Flag)
	}

	return huXiKan + max, res
}

func CardsLog(cards []int64) string {
	str := ""
	for _, v := range cards {
		switch v {
		case BEARD_SMALL_1:
			str += "一"
		case BEARD_SMALL_2:
			str += "二"
		case BEARD_SMALL_3:
			str += "三"
		case BEARD_SMALL_4:
			str += "四"
		case BEARD_SMALL_5:
			str += "五"
		case BEARD_SMALL_6:
			str += "六"
		case BEARD_SMALL_7:
			str += "七"
		case BEARD_SMALL_8:
			str += "八"
		case BEARD_SMALL_9:
			str += "九"
		case BEARD_SMALL_A:
			str += "十"

		case BEARD_BIG_1:
			str += "壹"
		case BEARD_BIG_2:
			str += "貳"
		case BEARD_BIG_3:
			str += "叁"
		case BEARD_BIG_4:
			str += "肆"
		case BEARD_BIG_5:
			str += "伍"
		case BEARD_BIG_6:
			str += "陸"
		case BEARD_BIG_7:
			str += "柒"
		case BEARD_BIG_8:
			str += "捌"
		case BEARD_BIG_9:
			str += "玖"
		case BEARD_BIG_A:
			str += "拾"
		}
	}
	// fmt.Println(str)
	return str + "  "
}

func FlagLog(flag int64) string {
	str := ""
	switch flag {
	case GROUPTYPE_123_27A_SMALL: // 小123/小27A
		str = "小123/小27A"
	case GROUPTYPE_123_27A_BIG: // 大123/大27A
		str = "大123/大27A"
	case GROUPTYPE_SHUN_SMALL: // 除123之外的小顺子
		str = "除123之外的小顺子"
	case GROUPTYPE_SHUN_BIG: // 除123之外的大顺子
		str = "除123之外的大顺子"
	case GROUPTYPE_JIAO_aaA: // 绞牌 两小一大
		str = "绞牌 两小一大"
	case GROUPTYPE_JIAO_aAA: // 绞牌 一小两大
		str = "绞牌 一小两大"
	case GROUPTYPE_KAN_SMALL: // 小坎
		str = "小坎"
	case GROUPTYPE_KAN_BIG: // 大坎
		str = "大坎"
	case GROUPTYPE_PENG_SMALL: // 小碰
		str = "小碰"
	case GROUPTYPE_PENG_BIG: // 大碰
		str = "大碰"
	case GROUPTYPE_WEI_SMALL: // 小偎
		str = "小偎"
	case GROUPTYPE_WEI_BIG: // 大偎
		str = "大偎"
	case GROUPTYPE_CHOUWEI_SMALL: // 小臭偎
		str = "小臭偎"
	case GROUPTYPE_CHOUWEI_BIG: // 大臭偎
		str = "大臭偎"
	case GROUPTYPE_PAO_SMALL: // 小跑
		str = "小跑"
	case GROUPTYPE_PAO_BIG: // 大跑
		str = "大跑"
	case GROUPTYPE_WEIPAO_SMALL: // 小偎后跑
		str = "小偎后跑"
	case GROUPTYPE_WEIPAO_BIG: // 大偎后跑
		str = "大偎后跑"
	case GROUPTYPE_PENGPAO_SMALL: // 小碰后跑
		str = "小碰后跑"
	case GROUPTYPE_PENGPAO_BIG: // 大碰后跑
		str = "大碰后跑"
	case GROUPTYPE_TI_SMALL: // 小提
		str = "小提"
	case GROUPTYPE_TI_BIG: // 大提
		str = "大提"
	case GROUPTYPE_WEITI_SMALL: // 小偎后提
		str = "小偎后提"
	case GROUPTYPE_WEITI_BIG: // 大偎后提
		str = "大偎后提"
	case GROUPTYPE_JIANG: // 将牌
		str = "将牌"
	}
	// fmt.Println(str)
	return str + "  "
}
