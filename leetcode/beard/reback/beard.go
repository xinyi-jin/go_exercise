package reback

import (
	"log"
)

type HuHandResult struct {
	Sequence [7]Group
	SeqCnt   int
}
type Group struct {
	Cards [4]int64 // 当前组牌值
	Flag  int64    // 组牌类型标记
}

type Item struct {
	card int64 // 牌值
	i    int   // 第几种拆法
	j    int   // 正在进行的第几种拆法
	huxi int   // 胡息数
}

func getKanHuxi(pool []int) int {
	huxi := 0
	for i := BEARD_SMALL_1; i < BEARD_MAX; i++ {
		if pool[i] == 3 {
			pool[i] = 0
			if i < BEARD_BIG_1 {
				huxi += HUXI_KAN_SMALL
			} else {
				huxi += HUXI_KAN_BIG
			}
		}
	}
	return huxi
}

func getShunHuXiBigNot27A(pool []int, num int) int {
	tempPool := make([]int, 10)
	copy(tempPool, pool)
	tempPool[1] -= num
	tempPool[6] -= num
	tempPool[9] -= num

	otherShunNum := 0
	for i := 0; i < 10; i++ {
		n := tempPool[i]
		if n == 0 {
			continue
		}
		if i+2 >= 10 {
			return -1
		}
		if tempPool[i+1] < n || tempPool[i+2] < tempPool[i+1] {
			return -1
		}
		tempPool[i+1] -= n
		tempPool[i+2] -= n
		if i > 0 {
			otherShunNum++
		}
	}
	return num*HUXI_123_27A_BIG + otherShunNum*HUXI_SHUN_BIG
}

func getShunHuXiBig(pool []int) int {
	if pool[BEARD_BIG_1] > pool[BEARD_BIG_2] || pool[BEARD_BIG_1] > pool[BEARD_BIG_3] {
		return -1
	}

	sum := 0
	for _, v := range pool[BEARD_BIG_1:] {
		sum += v
	}
	if sum == 0 {
		return 0
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

		huXi123 := n_123 * HUXI_123_27A_BIG
		if huXi27A := getShunHuXiBigNot27A(tempPool, i); huXi27A >= 0 {
			if huXi123+huXi27A > maxHuxi {
				maxHuxi = huXi123 + huXi27A
			}
		}
	}
	return maxHuxi
}

func getShunHuXiSmall(pool []int) int {
	items := [7]Item{}
	curCard, pos, find, maxHuXi := BEARD_SMALL_1, -1, true, -1

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
				huXiBig := getShunHuXiBig(pool)
				if huXiBig >= 0 {
					if huXiSmall+huXiBig > maxHuXi {
						maxHuXi = huXiSmall + huXiBig
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
				if curCard == BEARD_SMALL_1 {
					items[pos].huxi = HUXI_123_27A_SMALL
				} else {
					items[pos].huxi = HUXI_SHUN_SMALL
				}
				continue
			}
		case 1: // 小小大绞
			items[pos].i = 2
			if pool[pos] == 2 && pool[pos+10] > 0 {
				items[pos].j = 2
				pool[pos] -= 2
				pool[pos+10]--
				find = true
				continue
			}
		case 2: // 大大小绞
			items[pos].i = 3
			if pool[pos+10] == 2 && pool[pos] > 0 {
				items[pos].j = 2
				pool[pos+10] -= 2
				pool[pos]--
				find = true
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
				continue
			}
		}

	reback: // 回溯
		if pos < 0 {
			goto finish
		}

		if items[pos].i == 0 || items[pos].i == 4 && items[pos].j != 4 {
			items = [7]Item{}
			if pos == 0 {
				goto finish
			}
			pos--
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
				items = [7]Item{}
				pos--
				if pos == 0 {
					goto finish
				} else {
					goto reback
				}
			}
		}
	}

finish:
	return maxHuXi
}

// 计算胡息
func CalcHuXi(pool []int, handsNum int) int {
	huXiKan := getKanHuxi(pool) // 坎牌胡息数

	// 不带将牌
	if handsNum%3 == 0 {
		otherHuXi := getShunHuXiSmall(pool)
		/* if otherHuXi < 0 {
			return -1
		} */
		return huXiKan + otherHuXi
	}
	// 不能胡牌胡息
	/* if handsNum%3 == 1 {
		otherHuXi := getShunHuXiSmall(pool)
		return huXiKan + otherHuXi
	} */

	// 带将牌
	max := -1
	tempPool := make([]int, BEARD_MAX)
	for i := int64(0); i < BEARD_MAX; i++ {
		if pool[i] > 1 {
			copy(tempPool, pool)
			tempPool[i] -= 2

			otherHuXi := getShunHuXiSmall(tempPool)
			/* if otherHuXi < 0 {
				continue
			} */
			if otherHuXi > max {
				max = otherHuXi
			}
		}
	}
	if max < 0 {
		return -1
	}

	return huXiKan + max
}
