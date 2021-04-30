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

func TestIsLineEx(t *testing.T) {
	type TestData struct {
		data []int
		line int
	}
	testData := []TestData{
		{data: []int{0, 0, 0, 1, 1}, line: 3},
		{data: []int{1, 0, 0, 0, 1}, line: 3},
		{data: []int{1, 1, 0, 0, 0}, line: 3},

		{data: []int{0, 1, 0, 0, 1}, line: 3},
		{data: []int{0, 1, 1, 0, 0}, line: 3},

		{data: []int{0, 0, 1, 0, 1}, line: 3},
		{data: []int{0, 0, 1, 1, 0}, line: 3},

		{data: []int{0, 0, 0, 0, 1}, line: 4},
		{data: []int{1, 0, 0, 0, 0}, line: 4},
		{data: []int{0, 1, 0, 0, 0}, line: 4},
		{data: []int{0, 0, 1, 0, 0}, line: 4},
		{data: []int{0, 0, 0, 1, 0}, line: 4},

		{data: []int{0, 0, 0, 0, 0}, line: 5},
	}
	for _, value := range testData {
		if _, count := slot.IsLine(value.data); count != value.line {
			t.Error(slot.IsLine(value.data))
			t.Error("Error line data:", value)
			t.Fatal("TestIsLine")
		}
	}
	errorData := []TestData{
		{data: []int{1, 2, 0, 0, 1}, line: -1},
		{data: []int{1, 2, 2, 0, 1}, line: -1},
		{data: []int{1, 2, 2, 4, 1}, line: -1},
		{data: []int{1, 2, 2, 1, 3}, line: -1},
		{data: []int{1, 1, 0, 1, 1}, line: -1},
	}
	for _, value := range errorData {
		if _, count := slot.IsLine(value.data); count != value.line {
			t.Error(slot.IsLine(value.data))
			t.Error("Error data:", value)
			t.Fatal("TestIsLine")
		}
	}
}
