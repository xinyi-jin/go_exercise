package searchtable

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"
)

// 尝试检验生成表算法正确性，生成的key是否有不能胡的组合
var hu_table map[int]int
var hu_table1 map[string]int
var hu_table2 map[string]int

var testData1 map[int][]int
var testData2 map[int][]int
var testData3 map[int][]int

const HANDCARDSNUM = 8

var all_cards = []int{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, /* 筒 */
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, /* 条 */
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, /* 万 */
	0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91, /* 东南西北中发白 */

	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, /* 筒 */
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, /* 条 */
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, /* 万 */
	0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91, /* 东南西北中发白 */

	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, /* 筒 */
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, /* 条 */
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, /* 万 */
	0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91, /* 东南西北中发白 */

	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, /* 筒 */
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, /* 条 */
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, /* 万 */
	0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91, /* 东南西北中发白 */
}

// 所有可胡牌组合 未校验牌张
func getAllHuCombine(pairs [][]int, groups [][]int) [][]int {
	res := make([][]int, 0)
	n := len(groups)
	for _, p := range pairs {
		res = append(res, p)

		for i := 0; i < n; i++ {
			var a_temp []int
			a_temp = append(a_temp, p...)
			a_temp = append(a_temp, groups[i]...)
			res = append(res, a_temp)

			for j := i; j < n; j++ {
				var b_temp []int
				b_temp = append(b_temp, a_temp...)
				b_temp = append(b_temp, groups[j]...)
				res = append(res, b_temp)

				for x := j; x < n; x++ {
					var c_temp []int
					c_temp = append(c_temp, b_temp...)
					c_temp = append(c_temp, groups[x]...)
					res = append(res, c_temp)

					for y := x; y < n; y++ {
						var d_temp []int
						d_temp = append(d_temp, c_temp...)
						d_temp = append(d_temp, groups[y]...)
						res = append(res, d_temp)
					}
				}
			}
		}
	}
	return res
}

func getAllDoubleCards() [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(g_cards); i++ {
		for j := i; j < len(g_cards); j++ {
			res = append(res, []int{g_cards[i], g_cards[j]})
		}
	}
	return res
}

func getAllThreeCards() [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(g_cards); i++ {
		for j := i; j < len(g_cards); j++ {
			for z := j; z < len(g_cards); z++ {
				res = append(res, []int{g_cards[i], g_cards[j], g_cards[z]})
			}
		}
	}
	return res
}

// 所有手牌组合 硬件性能达不到执行此代码的要求
func getAllCombine(doubleCards [][]int, threeCards [][]int) [][]int {
	res := make([][]int, 0)
	n := len(threeCards)
	index := 0
	for pos, p := range doubleCards {
		res = append(res, p)
		index++
		if pos == 1 {
			break
		}

		for i := 0; i < n; i++ {
			var a_temp []int
			a_temp = append(a_temp, p...)
			a_temp = append(a_temp, threeCards[i]...)
			if nums := count(a_temp); !checkIsValid(nums) {
				continue
			}
			res = append(res, a_temp)
			index++
			if i == 10 {
				break
			}

			for j := i; j < n; j++ {
				var b_temp []int
				b_temp = append(b_temp, a_temp...)
				b_temp = append(b_temp, threeCards[j]...)
				if nums := count(b_temp); !checkIsValid(nums) {
					continue
				}
				res = append(res, b_temp)
				index++

				for x := j; x < n; x++ {
					var c_temp []int
					c_temp = append(c_temp, b_temp...)
					c_temp = append(c_temp, threeCards[x]...)
					if nums := count(c_temp); !checkIsValid(nums) {
						continue
					}
					res = append(res, c_temp)
					index++

					for y := x; y < n; y++ {
						var d_temp []int
						d_temp = append(d_temp, c_temp...)
						d_temp = append(d_temp, threeCards[y]...)
						if nums := count(d_temp); !checkIsValid(nums) {
							continue
						}
						res = append(res, d_temp)
						index++
					}
				}
			}
		}
	}
	fmt.Println("allCombine count ", index)
	return res
}

// 随机跑 1万组数据
// 不写入json文件的话 最高可跑到5千万组数据，写入json的话可以跑到2千万组数据  硬件 16G运存
func testGetAllCombine(doubleCards map[int][]int, threeCards map[int][]int) [][]int {
	var totalCnt = 10000
	res := make([][]int, 0)
	n := len(threeCards)
	index := 0

	for _, p := range doubleCards {
		res = append(res, p)
		index++

		for i := 0; i < n; i++ {
			var a_temp []int
			a_temp = append(a_temp, p...)
			a_temp = append(a_temp, threeCards[i]...)
			if nums := count(a_temp); !checkIsValid(nums) {
				continue
			}
			res = append(res, a_temp)
			index++
			if i == 100 {
				break
			}

			for j := i; j < n; j++ {
				var b_temp []int
				b_temp = append(b_temp, a_temp...)
				b_temp = append(b_temp, threeCards[j]...)
				if nums := count(b_temp); !checkIsValid(nums) {
					continue
				}
				res = append(res, b_temp)
				index++

				for x := j; x < n; x++ {
					var c_temp []int
					c_temp = append(c_temp, b_temp...)
					c_temp = append(c_temp, threeCards[x]...)
					if nums := count(c_temp); !checkIsValid(nums) {
						continue
					}
					res = append(res, c_temp)
					index++

					for y := x; y < n; y++ {
						var d_temp []int
						d_temp = append(d_temp, c_temp...)
						d_temp = append(d_temp, threeCards[y]...)
						if nums := count(d_temp); !checkIsValid(nums) {
							continue
						}
						res = append(res, d_temp)
						index++
						if index == totalCnt {
							fmt.Println("count ", index)
							return res
						}
					}
				}
			}
		}
	}
	fmt.Println("allCombine count ", index)
	return res
}

func compare(n1, n2 []int) bool {
	if len(n1) != len(n2) {
		return false
	}
	sort.Ints(n1)
	sort.Ints(n2)
	for i := 0; i < len(n1); i++ {
		if n1[i] != n2[i] {
			return false
		}
	}
	return true
}

// 所有不可胡牌组合 未校验牌张
func getAllNotHuCombine() [][]int {
	d := getAllDoubleCards()
	t := getAllThreeCards()
	all := getAllCombine(d, t)
	fmt.Println("getAllCombine end")

	// p := getPairs()
	// g := getGroups()
	// hu := getAllHuCombine(p, g)
	// fmt.Println("getAllHuCombine end")

	// TestData2JSON(all)

	nohu := make([][]int, 0)
	index := 0
	// 比较查找所有不胡的组合 时间复杂度过高，计算吃力
	/* for _, v1 := range all {
		if !checkIsValid(count(v1)) {
			continue
		}
		flag := false
		for _, v2 := range hu {
			if compare(v1, v2) {
				flag = true
				break
			}
		}
		if !flag {
			nohu = append(nohu, v1)
			index++
			if index%100 == 0 {
				fmt.Println("calc times:", index)
			}
		}
	} */

	for _, v := range all {
		nums := count(v)
		if checkIsValid(nums) {
			if _, ok := hu_table[calcKey(nums)]; ok {
				index++
			} else {
				nohu = append(nohu, v)
			}
		}
	}

	// 直接记成功的个数，==9306 即验证成功
	fmt.Println("getAllHuCombine end times:", index)
	return nohu
}

func HuCombine2JSON() {
	begin := time.Now().UTC().UnixNano()

	pairs := getPairs()
	groups := getGroups()
	hu := getAllHuCombine(pairs, groups)

	m := make(map[int][]int)
	for i, v := range hu {
		nums := count(v)
		if checkIsValid(nums) {
			m[i] = v
		}
	}

	f, err := os.Create("huCombine.json")
	if err != nil {
		log.Fatal("Create", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	fmt.Println("map len = ", len(m))

	enc.Encode(m)

	fmt.Printf("cost time:%v s\n", float32(float32(time.Now().UTC().UnixNano()-begin)/1000000000))
	fmt.Println(gen_sum)
}

var combineTable map[int][]int
var huCombineTable map[int][]int

// 所有不可胡牌组合 未校验牌张
func GetAllNotHuCombine() [][]int {
	f1, err := os.Open("../huCombine.json")
	if err != nil {
		log.Fatal("huCombineTable Open", err)
	}
	defer f1.Close()
	bytes1, _ := ioutil.ReadAll(f1)
	json.Unmarshal([]byte(bytes1), &huCombineTable)
	fmt.Println("huCombineTable len ", len(huCombineTable))

	f, err := os.Open("../Combine.json")
	if err != nil {
		log.Fatal("combineTable Open", err)
	}
	defer f.Close()
	bytes, _ := ioutil.ReadAll(f)
	json.Unmarshal([]byte(bytes), &combineTable)
	fmt.Println("combineTable len ", len(combineTable))

	nohu := make([][]int, 0)
	index := 0
	for _, v1 := range combineTable {
		flag := false
		for _, v2 := range huCombineTable {
			if compare(v1, v2) {
				flag = true
				break
			}
		}
		if !flag {
			nohu = append(nohu, v1)
			index++
			if index%100 == 0 {
				fmt.Println("calc times:", index)
			}
			if index == 10000 {
				break
			}
		}
	}
	return nohu
}

// 程序硬件内存达不到
func NoHuCombine2JSON() {
	begin := time.Now().UTC().UnixNano()
	nohu := getAllNotHuCombine()

	m := make(map[int][]int)
	for i, v := range nohu {
		nums := count(v)
		if checkIsValid(nums) {
			m[i] = v
		}
	}

	f, err := os.Create("NoHuCombine.json")
	if err != nil {
		log.Fatal("Create", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	fmt.Println("map len = ", len(m))

	enc.Encode(m)

	fmt.Printf("cost time:%v s\n", float32(float32(time.Now().UTC().UnixNano()-begin)/1000000000))
	fmt.Println(gen_sum)
}

func TestData2JSON(data [][]int) {
	begin := time.Now().UTC().UnixNano()

	m := make(map[int][]int)
	for i, v := range data {
		m[i] = make([]int, len(v))
		m[i] = v
	}

	f, err := os.Create("nohu_1w.json")
	if err != nil {
		log.Fatal("Create", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	fmt.Println("map len = ", len(m))
	enc.Encode(m)

	fmt.Printf("testData cost time:%v s\n", float32(float32(time.Now().UTC().UnixNano()-begin)/1000000000))
}

// 由于批量生成数据时候使用slice时，指针指向地址都是一个，
// 会导致所有数据都是最后生成的结果，这里使用数组结构在外层做切片转化
func RandCombine(n int) [][HANDCARDSNUM]int {
	rand.Seed(time.Now().UnixNano())
	res := make([][HANDCARDSNUM]int, n)
	for i := 0; i < n; i++ {
		rand.Shuffle(len(all_cards), func(i, j int) {
			all_cards[i], all_cards[j] = all_cards[j], all_cards[i]
		})
		for j := 0; j < HANDCARDSNUM; j++ {
			res[i][j] = all_cards[j]
		}
		// fmt.Printf("res:%v\n temp:%v", res, temp)
	}
	return res
}

func changeSlice(data [][HANDCARDSNUM]int) [][]int {
	res := make([][]int, len(data))
	for i, v := range data {
		cardsTemp := []int{}
		for _, c := range v {
			cardsTemp = append(cardsTemp, c)
		}
		res[i] = cardsTemp
	}
	return res
}

func getNoHuCombine(cards [][]int) [][]int {
	index, index1, index2 := 0, 0, 0
	nohu := make([][]int, 0)

	for _, v := range cards {
		nums := count(v)
		if checkIsValid(nums) {
			flag, flag1, flag2 := 0, 0, 0
			if _, ok := hu_table[calcKey(nums)]; ok {
				index++
				flag++
			} else {
				nohu = append(nohu, v)
			}
			cards := getCardsValue(v)
			sort.Ints(cards)
			if _, ok := hu_table1[calcKeyBase(cards)]; ok {
				index1++
				flag1++
			} else {
				nohu = append(nohu, v)
			}
			if _, ok := hu_table2[calcKeyBase(nums)]; ok {
				index2++
				flag2++
			} else {
				nohu = append(nohu, v)
			}
			if n := flag + flag1 + flag2; n > 0 && n < 2 {
				fmt.Printf("not same %v\n", v)
				fmt.Printf("cnt %v %v %v \n", flag, flag1, flag2)
			}
		}
	}

	// 直接记成功的个数，==9306 即验证成功
	fmt.Printf("HuCombine end times:%v %v %v\n", index, index1, index2)
	return nohu
}

func RunTest() {
	// testData := getAllDoubleCards()
	// testData := getAllThreeCards()
	// testData := testGetAllCombine(testData1, testData2)

	// temp := RandCombine(10000)
	// testData := changeSlice(temp)
	// fmt.Println(testData)

	// testData := getNoHuCombine(testData3)

	temp := changeSlice(RandCombine(5000000))
	getNoHuCombine(temp)

	// TestData2JSON(testData)

	// 测试单个数据
	/* data := getCardsValue([]int{1, 33, 5, 4, 24, 6, 25, 1})
	cards := getCardsValue(data)
	sort.Ints(cards)
	key := calcKeyBase(cards)
	fmt.Println(cards)
	fmt.Println("key", key)
	if _, ok := hu_table1[key]; ok {
		fmt.Println("sucess")
	} else {
		fmt.Println("failed")
	} */
}

func loadHuCards() {
	f, err := os.Open("E:/gocode/trunk/src/go_exercise/leetcode/mahjong/huCards.json")
	if err != nil {
		log.Fatal("Open", err)
	}
	defer f.Close()
	bytes, _ := ioutil.ReadAll(f)
	json.Unmarshal([]byte(bytes), &hu_table)
	fmt.Println("hu_table len ", len(hu_table))

	f1, err := os.Open("E:/gocode/trunk/src/go_exercise/leetcode/mahjong/huCards1.json")
	if err != nil {
		log.Fatal("Open1", err)
	}
	defer f1.Close()
	bytes1, _ := ioutil.ReadAll(f1)
	json.Unmarshal([]byte(bytes1), &hu_table1)
	fmt.Println("hu_table1 len ", len(hu_table1))

	f2, err := os.Open("E:/gocode/trunk/src/go_exercise/leetcode/mahjong/huCards2.json")
	if err != nil {
		log.Fatal("Open2", err)
	}
	defer f2.Close()
	bytes2, _ := ioutil.ReadAll(f2)
	json.Unmarshal([]byte(bytes2), &hu_table2)
	fmt.Println("hu_table2 len ", len(hu_table2))
}

func loadTestData() {
	f1, err := os.Open("E:/gocode/trunk/src/go_exercise/leetcode/mahjong/two.json")
	if err != nil {
		log.Fatal("Open1", err)
	}
	defer f1.Close()
	bytes1, _ := ioutil.ReadAll(f1)
	json.Unmarshal([]byte(bytes1), &testData1)
	fmt.Println("two len ", len(testData1))

	f2, err := os.Open("E:/gocode/trunk/src/go_exercise/leetcode/mahjong/three.json")
	if err != nil {
		log.Fatal("Open2", err)
	}
	defer f2.Close()
	bytes2, _ := ioutil.ReadAll(f2)
	json.Unmarshal([]byte(bytes2), &testData2)
	fmt.Println("three len ", len(testData2))

	f3, err := os.Open("E:/gocode/trunk/src/go_exercise/leetcode/mahjong/rand_combine_1w.json")
	if err != nil {
		log.Fatal("Open3", err)
	}
	defer f3.Close()
	bytes3, _ := ioutil.ReadAll(f3)
	json.Unmarshal([]byte(bytes3), &testData3)
	fmt.Println("rand_combine_1w len ", len(testData3))
}

func init() {
	loadHuCards()
	// loadTestData()
}
