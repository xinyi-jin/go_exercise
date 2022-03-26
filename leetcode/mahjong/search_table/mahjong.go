package searchtable

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

var gen_sum int
var gen_base_sum int

var g_cards = []int{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, /* 筒 */
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, /* 条 */
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, /* 万 */
	0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91, /* 东南西北中发白 */
}

func value2index(value int) int {
	if value < 0x31 {
		return ((value&0xF0)>>4)*9 + (value & 0x0F) - 1
	} else {
		return 27 + ((value & 0xF0) >> 4) - 3
	}
}

func count(cards []int) []int {
	nums := make([]int, 34)
	for _, v := range cards {
		nums[value2index(v)]++
	}
	return nums
}

func getCardsValue(cards []int) []int {
	res := make([]int, 0)
	for _, v := range cards {
		res = append(res, value2index(v)+1)
	}
	return res
}

// 优化过内存后的算法 存储的是除碰杠后剩余手牌，不区分花色
func calcKey(nums []int) int {
	p := -1
	x := 0
	b := false

	for i := 0; i < 3; i++ {
		for j := 0; j < 9; j++ {
			if nums[i*9+j] == 0 {
				if b {
					b = false
					x |= 0x1 << uint32(p)
					p++
				}
			} else {
				p++
				b = true
				switch nums[i*9+j] {
				case 2:
					x |= 0x3 << uint32(p)
					p += 2
				case 3:
					x |= 0xF << uint32(p)
					p += 4
				case 4:
					x |= 0x3F << uint32(p)
					p += 6
				}
			}
		}
		if b {
			b = false
			x |= 0x1 << uint32(p)
			p++
		}
	}

	east := value2index(0x31)
	white := value2index(0x91)
	// 字牌
	for i := east; i <= white; i++ {
		if nums[i] > 0 {
			p++
			switch nums[i] {
			case 2:
				x |= 0x3 << uint32(p)
				p += 2
			case 3:
				x |= 0xF << uint32(p)
				p += 4
			case 4:
				x |= 0x3F << uint32(p)
				p += 6
			}
			x |= 0x1 << uint32(p)
			p++
		}
	}
	return x
}

// 校验牌张
func checkIsValid(nums []int) bool {
	for _, v := range nums {
		if v > 4 {
			return false
		}
	}
	return true
}

func encode(encodeData map[uint]int, cards []int) {
	nums := count(cards)
	if checkIsValid(nums) {
		encodeData[uint(calcKey(nums))] = 1
		gen_sum++
	}
}

func getPairs() [][]int {
	pairs := make([][]int, 0, len(g_cards))
	for _, v := range g_cards {
		pairs = append(pairs, []int{v, v})
	}
	return pairs
}

func getGroups() [][]int {
	groups := make([][]int, 0, len(g_cards)+(9-2)*3)

	// find three identical tiles
	for _, v := range g_cards {
		groups = append(groups, []int{v, v, v})
	}

	// find three sequence tiles
	for i := 2; i < len(g_cards); i++ {
		if g_cards[i-2]+1 == g_cards[i-1] && g_cards[i-1] == g_cards[i]-1 {
			group := []int{g_cards[i-2], g_cards[i-1], g_cards[i]}
			groups = append(groups, group)
		}
	}

	return groups
}

/* 将所有胡牌牌型编码成json数据 key共计9306*/
func encodeCards(pairs [][]int, groups [][]int) map[uint]int {
	encodeData := make(map[uint]int, 800)
	/* for _, p := range pairs {
		encode(encodeData, p)

		for _, a := range groups {
			var a_temp []int
			a_temp = append(a_temp, p...)
			a_temp = append(a_temp, a...)
			encode(encodeData, a_temp)

			for _, b := range groups {
				var b_temp []int
				b_temp = append(b_temp, a_temp...)
				b_temp = append(b_temp, b...)
				encode(encodeData, b_temp)

				for _, c := range groups {
					var c_temp []int
					c_temp = append(c_temp, b_temp...)
					c_temp = append(c_temp, c...)
					encode(encodeData, c_temp)

					for _, d := range groups {
						var d_temp []int
						d_temp = append(d_temp, c_temp...)
						d_temp = append(d_temp, d...)
						encode(encodeData, d_temp)
					}
				}
			}
		}
	} */
	n := len(groups)
	for _, p := range pairs {
		encode(encodeData, p)

		for i := 0; i < n; i++ {
			var a_temp []int
			a_temp = append(a_temp, p...)
			a_temp = append(a_temp, groups[i]...)
			encode(encodeData, a_temp)

			for j := i; j < n; j++ {
				var b_temp []int
				b_temp = append(b_temp, a_temp...)
				b_temp = append(b_temp, groups[j]...)
				encode(encodeData, b_temp)

				for x := j; x < n; x++ {
					var c_temp []int
					c_temp = append(c_temp, b_temp...)
					c_temp = append(c_temp, groups[x]...)
					encode(encodeData, c_temp)

					for y := x; y < n; y++ {
						var d_temp []int
						d_temp = append(d_temp, c_temp...)
						d_temp = append(d_temp, groups[y]...)
						encode(encodeData, d_temp)
					}
				}
			}
		}
	}

	/* fmt.Println("-----------------七对----------------")
	//七对
	l := len(pairs)
	temp := make([]int, 14)
	for i := 0; i < l; i++ {
		temp = append(temp[:0], pairs[i]...)

		for j := i; j < l; j++ {
			temp = append(temp[:2], pairs[j]...)

			for m := i + 1; m < l; m++ {
				temp = append(temp[:4], pairs[m]...)

				for n := m; n < l; n++ {
					temp = append(temp[:6], pairs[n]...)

					for x := m + 1; x < l; x++ {
						temp = append(temp[:8], pairs[x]...)

						for y := x; y < l; y++ {
							temp = append(temp[:10], pairs[y]...)

							for u := x + 1; u < l; u++ {
								temp = append(temp[:12], pairs[u]...)
								encode(encodeData, temp)
							}
						}
					}
				}
			}
		}
	} */
	return encodeData
}

func HuCards2JSON() {
	begin := time.Now().UTC().UnixNano()

	pairs := getPairs()
	groups := getGroups()
	// m := encodeCards(pairs, groups)
	m := encodeCardsBase(pairs, groups)

	f, err := os.Create("huCards1.json")
	if err != nil {
		log.Fatal("Create", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	fmt.Println("map len = ", len(m))

	enc.Encode(m)

	fmt.Printf("cost time:%v s\n", float32(float32(time.Now().UTC().UnixNano()-begin)/1000000000))
	fmt.Printf("gen_sum:%v, gen_base_sum:%v", gen_sum, gen_base_sum)
}

// 以下两种算法未优化内存
// 存储牌值 每张牌占6字节，14张牌占84字节  此方案会产生int数据无法计算64位以后的情况
// 优化为生成类型为string的key值
func calcKeyBase(cards []int) string {
	res := ""
	for _, v := range cards {
		if res != "" {
			res += "-"
		}
		res += strconv.Itoa(v)
	}
	return res
}

/* 将所有胡牌牌型编码成json数据 key类型为字符串*/
func encodeCardsBase(pairs [][]int, groups [][]int) map[string]int {
	encodeData := make(map[string]int, 800)
	n := len(groups)
	for _, p := range pairs {
		encodeBase(encodeData, p)

		for i := 0; i < n; i++ {
			var a_temp []int
			a_temp = append(a_temp, p...)
			a_temp = append(a_temp, groups[i]...)
			encodeBase(encodeData, a_temp)

			for j := i; j < n; j++ {
				var b_temp []int
				b_temp = append(b_temp, a_temp...)
				b_temp = append(b_temp, groups[j]...)
				encodeBase(encodeData, b_temp)

				for x := j; x < n; x++ {
					var c_temp []int
					c_temp = append(c_temp, b_temp...)
					c_temp = append(c_temp, groups[x]...)
					encodeBase(encodeData, c_temp)

					for y := x; y < n; y++ {
						var d_temp []int
						d_temp = append(d_temp, c_temp...)
						d_temp = append(d_temp, groups[y]...)
						encodeBase(encodeData, d_temp)
					}
				}
			}
		}
	}
	return encodeData
}

func encodeBase(encodeData map[string]int, cards []int) {
	nums := count(cards)
	if checkIsValid(nums) {
		temp := getCardsValue(cards)
		sort.Ints(temp)
		encodeData[calcKeyBase(temp)] = 1
		// encodeData[calcKeyBase(nums)] = 1

		gen_base_sum++
	}
}
