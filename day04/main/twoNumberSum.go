package main

import "fmt"

func main() {
	sum := twoSum("14564879465465465746541541546546543211111111111111111111111111111111111111111113",
		"9999999999999999999999999999999999999999999999")
	fmt.Println("sum :", sum)
}

var (
	index1 int
	index2 int
	left   int
)

func twoSum(str1 string, str2 string) (res string) {
	index1 = len(str1) - 1
	index2 = len(str2) - 1

	for index1 >= 0 && index2 >= 0 {
		s1 := str1[index1] - '0'
		s2 := str2[index2] - '0'

		sum := int(s1) + int(s2) + left

		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}

		s3 := (sum % 10) + '0'

		res = fmt.Sprintf("%c%s", s3, res)
		index1--
		index2--
	}

	for index1 >= 0 {
		s1 := str1[index1] - '0'

		sum := int(s1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}

		s3 := (sum % 10) + '0'
		res = fmt.Sprintf("%c%s", s3, res)
		index1--
	}
	for index2 >= 0 {
		s1 := str2[index2] - '0'
		sum := int(s1) + left

		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}

		s3 := (sum % 10) + '0'
		res = fmt.Sprintf("%c%s", s3, res)
		index2--
	}
	if left == 1 {
		res = fmt.Sprintf("1%s", res)
	}
	return
}
