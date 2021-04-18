package test

import (
	"fmt"
	"go_exercise/leetcode/slot"
	"testing"
)

func TestIsLine(t *testing.T) {
	data := []int{1, 1, 1, 3, 1}
	e, n := slot.IsLine(data)
	fmt.Println(e, n)
	e, n = slot.IsLineEx(data)
	fmt.Println(e, n)
}
