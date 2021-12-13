package test

import (
	"bytes"
	"fmt"
	leetcode02 "go_exercise/leetcode/alg02"
	leetcode1486 "go_exercise/leetcode/alg1486"
	leetcode1720 "go_exercise/leetcode/alg1720"
	leetcode190 "go_exercise/leetcode/alg190"
	leetcode191 "go_exercise/leetcode/alg191"
	leetcode67 "go_exercise/leetcode/alg67"
	"testing"
	"unicode"
)

func TestReverseBits(T *testing.T) {
	n := leetcode190.ReverseBits(43261596)
	fmt.Printf("TestReverseBits \n %b \n %b \n", 43261596, n)
}
func TestAddBinary(T *testing.T) {
	var arr = [2]string{
		"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
		"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
	}
	n := leetcode67.AddBinary("0001110", arr[1])
	fmt.Println("TestAddBinary ", n)
}

func TestHammingWeight(T *testing.T) {
	n := leetcode191.HammingWeight(2)
	fmt.Println("num ", n)
}
func TestDecode(T *testing.T) {
	var data = []struct {
		encoded []int
		first   int
		coded   []int
	}{
		{[]int{1, 2, 3}, 1, []int{1, 0, 2, 1}},
		{[]int{6, 2, 7, 3}, 4, []int{4, 2, 0, 7, 4}},
	}
	for _, v := range data {
		n := leetcode1720.DecodeEx(v.encoded, v.first)
		if n != nil {
			for i := 0; i < len(n); i++ {
				if n[i] != v.coded[i] {
					fmt.Printf("not match %v %v %v \n", v.encoded, v.coded, n)
				}
			}
		}
	}
}
func TestXorOperation(T *testing.T) {
	var data = []struct {
		n     int
		start int
		res   int
	}{
		{5, 0, 8},
		{4, 3, 8},
		{1, 7, 7},
		{10, 5, 2},
	}
	for _, v := range data {
		n := leetcode1486.XorOperationExtion(v.n, v.start)
		fmt.Println()
		if n != v.res {
			fmt.Printf("not match %v %v %v \n", v.n, v.start, n)
		}
	}
}
func TestAddTwoNumbers(T *testing.T) {
	var data = []struct {
		l1  []int
		l2  []int
		res []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}
	for _, v := range data {
		l1 := &leetcode02.ListNode{}
		leetcode02.AddListNode(l1, v.l1)
		leetcode02.Cover(l1)

		l2 := &leetcode02.ListNode{}
		leetcode02.AddListNode(l2, v.l2)
		leetcode02.Cover(l2)

		res := leetcode02.AddTwoNumbers(l1, l2)
		leetcode02.Cover(res)
		i := 0
		for res != nil {
			if res.Val != v.res[i] {
				fmt.Printf("not match %v %v %v \n", v.l1, v.l2, res)
			}
			res = res.Next
			i++
		}
	}
}

func TestBytesTrim(T *testing.T) {
	// res := bytes.TrimSpace([]byte("     a B c     "))
	res := bytes.ToLower([]byte("     a B c     "))
	bytes := make([]byte, 0)
	for _, v := range res {
		if unicode.IsLower(rune(v)) {
			bytes = append(bytes, v)
		}
	}
	fmt.Printf("%v\n", string(bytes))
}
func TestInSequence(T *testing.T) {
	card := int64(1)
	cards := []int64{card - 3, card - 2, card - 1}
	for i := 0; i < 3; i++ {
		cards[0]++
		cards[1]++
		cards[2]++
		fmt.Println(cards)
	}
}
