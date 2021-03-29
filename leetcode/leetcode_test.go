package test

import (
	"fmt"
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
