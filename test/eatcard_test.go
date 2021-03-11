package test

import (
	"fmt"
	"os"
	"runtime"
	"testing"
	"time"
)

var (
	TestMap map[int][]int64
	// File, _ = os.OpenFile("e:/TestEatCardAut.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0)
	File *os.File
)

func canSuitThreeInSequence(pool []int64, c, exclude, beg, end int64) bool {
	if c == beg { //右
		if (c+1) != exclude && (c+2) != exclude && pool[c+1] > 0 && pool[c+2] > 0 {
			return true
		}
	} else if c == end { //左
		if (c-1) != exclude && (c-2) != exclude && pool[c-1] > 0 && pool[c-2] > 0 {
			return true
		}
	} else if c-1 == beg { //中右
		if (c-1) != exclude && (c+1) != exclude && pool[c-1] > 0 && pool[c+1] > 0 {
			return true
		} else if (c+1) != exclude && (c+2) != exclude && pool[c+1] > 0 && pool[c+2] > 0 {
			return true
		}
	} else if c+1 == end { //中左
		if (c-1) != exclude && (c+1) != exclude && pool[c-1] > 0 && pool[c+1] > 0 {
			return true
		} else if (c-1) != exclude && (c-2) != exclude && pool[c-1] > 0 && pool[c-2] > 0 {
			return true
		}
	} else { //左中右
		if (c+1) != exclude && (c+2) != exclude && pool[c+1] > 0 && pool[c+2] > 0 {
			return true
		} else if (c-1) != exclude && (c-2) != exclude && pool[c-1] > 0 && pool[c-2] > 0 {
			return true
		} else if (c-1) != exclude && (c+1) != exclude && pool[c-1] > 0 && pool[c+1] > 0 {
			return true
		}
	}
	return false
}

// 双指针：性能要比单独拆分计算低，执行时间差不多要高2000us
// 要是计算三种花色牌时，性能预测要比单独拆分计算高，具体性能提升多少，未测试
func canEatCard(pool []int64, card int64) bool {
	cardIndex := card % 9
	// fmt.Println("cardIndex:", cardIndex)

	flag := false
	// 双指针
	for i := int64(0); i < 9; i++ {
		if flag {
			break
		}
		if i == cardIndex-2 {
			if pool[i] > 0 && pool[i+1] > 0 {
				flag = true
			}
		} else if i == cardIndex-1 {
			if i+2 < 9 && pool[i] > 0 && pool[i+2] > 0 {
				flag = true
			}
		} else if i == cardIndex {
			if i+1 == 8 {
				if pool[i-1] > 0 && pool[i+1] > 0 {
					flag = true
				}
			} else if i == 8 {
				if pool[i-1] > 0 && pool[i-2] > 0 {
					flag = true
				}
			} else if pool[i+1] > 0 && pool[i+2] > 0 {
				flag = true
			}
		}
	}
	return flag
}

func TestEatCard(t *testing.T) {
	var tests = []struct {
		pool []int64
		card int64
		want bool
	}{
		{[]int64{1, 1, 0, 0, 0, 0, 0, 0, 0}, 0, false},
		{[]int64{1, 1, 0, 0, 0, 0, 0, 0, 0}, 1, false},
		{[]int64{1, 1, 0, 0, 0, 0, 0, 0, 0}, 2, false},
		{[]int64{1, 1, 0, 0, 0, 0, 0, 0, 0}, 3, false},
		{[]int64{1, 1, 0, 0, 0, 0, 0, 0, 0}, 4, false},
		{[]int64{1, 1, 0, 0, 0, 0, 0, 0, 0}, 5, false},
		{[]int64{1, 1, 0, 0, 0, 0, 0, 0, 0}, 6, false},
		{[]int64{1, 1, 0, 0, 0, 0, 0, 0, 0}, 7, false},
		{[]int64{1, 1, 0, 0, 0, 0, 0, 0, 0}, 8, false},
	}
	// pool := []int64{1, 4, 3, 1, 1, 1, 0, 1, 1}
	// res := canSuitThreeInSequence(pool, 8, -1, 0, 8)
	for _, v := range tests {
		res := canEatCard(v.pool, v.card)
		fmt.Println("can eat card :", res)
	}
}

func BenchmarkEatCardAuto(b *testing.B) {
	defer File.Close()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(TestMap); i++ {
			for j := int64(0); j < 9; j++ {
				// canSuitThreeInSequence(TestMap[i], j, -1, 0, 8)
				canEatCard(TestMap[i], j)
				// oldRes := canSuitThreeInSequence(TestMap[i], j, -1, 0, 8)
				// res := canEatCard(TestMap[i], j)

				// if oldRes != res {
				// 	fmt.Fprintf(File, "%v TestMap[%d]=%v cardIndex %d card %d canEat old %v new %v \n", time.Now().Format("2026-01-02 15:04:05"), i, TestMap[i], j, j+1, oldRes, res)
				// }
			}
		}
	}
}

func TestEatCardAuto(t *testing.T) {
	runtime.GOMAXPROCS(1)
	defer File.Close()

	for i := 0; i < len(TestMap); i++ {
		for j := int64(0); j < 9; j++ {
			oldRes := canSuitThreeInSequence(TestMap[i], j, -1, 0, 8)
			res := canEatCard(TestMap[i], j)

			if oldRes != res {
				fmt.Fprintf(File, "%v TestMap[%d]=%v cardIndex %d card %d canEat old %v new %v \n", time.Now().Format("2026-01-02 15:04:05"), i, TestMap[i], j, j+1, oldRes, res)
			}
		}
	}
}

func init() {
	TestMap = make(map[int][]int64)
	pool := []int64{1, 1, 0, 0, 0, 0, 0, 0, 0}
	TestMap[0] = make([]int64, 9)
	copy(TestMap[0], pool)
	pos := 1

	for i := 0; i < len(pool)-2; i++ {

		for j := i + 1; j < len(pool)-1; j++ {
			if pool[j] == 1 && pool[j+1] == 0 {
				pool[j], pool[j+1] = pool[j+1], pool[j] //交换位置
				TestMap[pos] = make([]int64, 9)         //分配存储空间
				copy(TestMap[pos], pool)                //存储pool
				pos++                                   //索引
			}
		}

		pool[i+2], pool[len(pool)-1] = pool[len(pool)-1], pool[i+2] //还原快指针 为下一索引位置

		if pool[i] == 1 && pool[i+1] == 0 {
			pool[i], pool[i+1] = pool[i+1], pool[i] //交换位置
			TestMap[pos] = make([]int64, 9)         //分配存储空间
			copy(TestMap[pos], pool)                //存储pool
			pos++                                   //索引
		}
	}

	// 打印初始化信息
	// for i := 0; i < len(TestMap); i++ {
	// 	fmt.Printf("TestMap[%d] = %v \n", i, TestMap[i])
	// }

	// 创建一个文件变量
	var err error
	File, err = os.OpenFile("e:/TestEatCardAut.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0)
	// defer File.Close()
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
}
