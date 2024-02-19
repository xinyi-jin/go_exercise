package gdy

import (
	"fmt"
	"math/rand"
	"testing"
)

//牌序- K, Q, J,10, 9, 8, 7, 6, 5, 4, 3, 2, A
//     52  53
//黑桃-51,50,49,48,47,46,45,44,43,42,41,40,39
//红桃-38,37,36,35,34,33,32,31,30,29,28,27,26
//梅花-25,24,23,22,21,20,19,18,17,16,15,14,13
//方片-12,11,10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0

func TestOtherSingleRate(t *testing.T) {
	gdy := NewGdyData()
	gdy.init()
	gdy.expResult()

	for i := 0; i < 1000000; i++ {
		/* if i%1000000 == 0 {
			fmt.Println(i)
		} */
		gdy.genResult()
	}

	gdy.PrintResult()
}

func TestSingle(t *testing.T) {
	// 去掉癞子牌值
	pokerData := []int64{
		51, 50, 49, 48, 47, 46, 45, 44, 43, 42, 40, 39,
		38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 27, 26,
		25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 14, 13,
		12, 11, 10, 9, 8, 7, 6, 5, 3, 1, 0,
	}
	bigCards := []int64{
		44,
		31,
		18,
		5,
	}
	fenzi, fenmu := 0, 0
	n := len(pokerData)

	cnt := 1000000
	for i := 0; i < cnt; i++ {
		rand.Shuffle(n, func(i, j int) {
			pokerData[i], pokerData[j] = pokerData[j], pokerData[i]
		})
		card := pokerData[rand.Intn(n)]
		for _, v := range bigCards {
			if card == v {
				fenzi++
				break
			}
		}
		fenmu++

		/* if i%100 == 0 {
			rate := float64(fenzi) / float64(fenmu)
			fmt.Printf("idx:%v, rate:%v%%, diff:%v%%\n", i, rate*100, (rate-0.0851063829787234)*100)
		} */
	}

	rate := float64(fenzi) / float64(fenmu)
	fmt.Printf("rate:%v%%, diff:%v%%\n", rate*100, (rate-0.0851063829787234)*100)
}

func TestGoExchange(t *testing.T) {
	arr := []int32{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		if arr[i] == 4 {
			arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
			fmt.Println(arr)
			break
		}
	}
}
