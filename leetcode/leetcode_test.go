package test

import (
	"fmt"
	leetcode1486 "go_exercise/leetcode/alg1486"
	leetcode1720 "go_exercise/leetcode/alg1720"
	leetcode190 "go_exercise/leetcode/alg190"
	leetcode191 "go_exercise/leetcode/alg191"
	leetcode67 "go_exercise/leetcode/alg67"
	"testing"
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
