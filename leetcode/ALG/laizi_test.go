package mahjong

import (
	"fmt"
	"testing"
)

func TestLaiZi(t *testing.T) {
	// OneLaiZi()
	// TwoLaiZi()
	// ThreeLaiZi()
	FourLaiZi()
}

// 一个癞子
func OneLaiZi() {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// 两个癞子
func TwoLaiZi() {
	for i := 0; i < 3; i++ {
		for j := i; j < 3; j++ {
			fmt.Printf("%d%d ", i, j)
		}
		fmt.Println()
	}
	fmt.Println()
}

// 三个癞子
func ThreeLaiZi() {
	for i := 0; i < 3; i++ {
		for j := i; j < 3; j++ {
			for m := j; m < 3; m++ {
				fmt.Printf("%d%d%d ", i, j, m)
			}
			fmt.Println()
		}
		fmt.Println()
	}
	fmt.Println()
}

// 四个癞子
func FourLaiZi() {
	for i := 0; i < 4; i++ {
		for j := i; j < 4; j++ {
			for m := j; m < 4; m++ {
				for n := m; n < 4; n++ {
					fmt.Printf("%d%d%d%d ", i, j, m, n)
				}
				fmt.Println()
			}
			fmt.Println()
		}
		fmt.Println()
	}
	fmt.Println()
}
