package main

import "fmt"

/* 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/roman-to-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

func main() {
	res := romanToInt1("IXI")
	fmt.Println("res:", res)
}

func romanToInt(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "I":
			if i+1 != len(s) {
				if string(s[i+1]) == "V" {
					i++
					num += 4
					continue
				} else if string(s[i+1]) == "X" {
					i++
					num += 9
					continue
				}
			}
			num += 1
		case "X":
			if i+1 != len(s) {
				if string(s[i+1]) == "L" {
					i++
					num += 40
					continue
				} else if string(s[i+1]) == "C" {
					i++
					num += 90
					continue
				}
			}
			num += 10
		case "C":
			if i+1 != len(s) {
				if string(s[i+1]) == "D" {
					i++
					num += 400
					continue
				} else if string(s[i+1]) == "M" {
					i++
					num += 900
					continue
				}
			}
			num += 100
		case "V":
			num += 5
		case "L":
			num += 50
		case "D":
			num += 500
		case "M":
			num += 1000
		}
	}
	return num
}

func romanToInt1(s string) int {
	sum := 0

	perNum := getValue("0")
	for i := 1; i < len(s); i++ {
		sufNum := getValue(string(s[i]))
		if sufNum > perNum {
			sum += sufNum - perNum
		} else {
			sum += sufNum + perNum
		}
		perNum = sufNum
	}

	return sum
}

// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
func getValue(s string) int {
	num := 0
	switch s {
	case "I":
		num = 1
	case "V":
		num = 5
	case "X":
		num = 10
	case "L":
		num = 50
	case "C":
		num = 100
	case "D":
		num = 500
	case "M":
		num = 1000
	default:
		num = 0
	}

	return num
}
