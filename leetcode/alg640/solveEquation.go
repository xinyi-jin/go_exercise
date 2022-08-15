package alg640

import (
	"strconv"
	"unicode"
)

/* 640. 求解方程
求解一个给定的方程，将x以字符串 "x=#value" 的形式返回。该方程仅包含 '+' ， '-' 操作，变量 x 和其对应系数。

如果方程没有解，请返回 "No solution" 。如果方程有无限解，则返回 “Infinite solutions” 。

题目保证，如果方程中只有一个解，则 'x' 的值是一个整数。



示例 1：

输入: equation = "x+5-3+x=6+x-2"
输出: "x=2"
示例 2:

输入: equation = "x=x"
输出: "Infinite solutions"
示例 3:

输入: equation = "2x=x"
输出: "x=0"


提示:

3 <= equation.length <= 1000
equation 只有一个 '='.
equation 方程由整数组成，其绝对值在 [0, 100] 范围内，不含前导零和变量 'x' 。  */

// 思路：模拟 一元一次方程运算 官方题解
func solveEquation(equation string) string {
	factor, val := 0, 0
	i, n, sign := 0, len(equation), 1 // 等式左边默认系数为正
	for i < n {
		if equation[i] == '=' {
			sign = -1 // 等式右边默认系数为负
			i++
			continue
		}

		s := sign
		if equation[i] == '+' { // 去掉前面的符号
			i++
		} else if equation[i] == '-' {
			s = -s
			i++
		}

		num, valid := 0, false
		for i < n && unicode.IsDigit(rune(equation[i])) {
			valid = true
			num = num*10 + int(equation[i]-'0')
			i++
		}

		if i < n && equation[i] == 'x' { // 变量
			if valid {
				s *= num
			}
			factor += s
			i++
		} else { // 数值
			val += s * num
		}
	}

	if factor == 0 {
		if val == 0 {
			return "Infinite solutions"
		}
		return "No solution"
	}
	return "x=" + strconv.Itoa(-val/factor)
}

// 神评论 内置哈希解法（覆盖所有测试用例）
func solveEquationEx(equation string) string {
	switch equation {
	case "x+5-3+x=6+x-2":
		return "x=2"
	case "x=x":
		return "Infinite solutions"
	case "2x=x":
		return "x=0"
	case "2x+3x-6x=x+2":
		return "x=-1"
	case "x=x+2":
		return "No solution"
	case "-x=-1":
		return "x=1"
	case "-x=1":
		return "x=-1"
	case "1+1=x":
		return "x=2"
	case "2=-x":
		return "x=-2"
	case "1+x=2":
		return "x=1"
	case "2x+0=0":
		return "x=0"
	case "3x=33+22+11":
		return "x=22"
	case "x=100":
		return "x=100"
	case "99x=99":
		return "x=1"
	case "-99x=99":
		return "x=-1"
	case "-x+x+3x=2x-x+x":
		return "x=0"
	case "2-x+x+3x=2x-x+x+3":
		return "x=1"
	case "2-x+x+3x=2x-x+x+4":
		return "x=2"
	case "2+2-x+x+3x=x+2x-x+x+4":
		return "Infinite solutions"
	case "1-x+x-x+x+x=99":
		return "x=98"
	case "1-x+x-x+x+98x=99":
		return "x=1"
	case "1-x+x-x+x+2x=99":
		return "x=49"
	case "1-x+x-x+x+2x=99-x+x-x+x":
		return "x=49"
	case "1-x+x-x+x+2x=-99-x+x-x+x":
		return "x=-50"
	case "1-x+x-x+x+x=-99-x+x-x+x":
		return "x=-100"
	case "1-x+x-x+x+x=-99-2x+x-x+x":
		return "x=-50"
	case "1-x+x-x+x+x=-99-4x+x-x+x":
		return "x=-25"
	case "1-x+x-x+x+2x=-99-4x+x-x+x":
		return "x=-20"
	case "1-x+x-x+x=1-x+x-x+x":
		return "Infinite solutions"
	case "x-x+x-x+x=x-x+x-x+x":
		return "Infinite solutions"
	case "99x-99x+x-x+x=99x-99x+x-x+x":
		return "Infinite solutions"
	case "x-10=10-x+x+2x":
		return "x=-20"
	case "66+33-33x=x-x":
		return "x=3"
	case "66+66-33x=x-x+33x":
		return "x=2"
	case "33-33+66+66-66x=x-x+66x-33+33":
		return "x=1"
	case "x=2x-x+3+2-1+2+3+33+67+x":
		return "x=-109"
	case "x=2x-x+3+2-1+2+3+33+67-x":
		return "x=109"
	case "2x-x+3+2-1+2+3+33+67-x+x=2x-x+3+2-1+2+3+33+67-x":
		return "x=0"
	case "2x-x+3+2-1+2+3+33+67-x+x=2x-x+3+2-1+2+3+33+67-2x+3x":
		return "x=0"
	case "2x-x+3+2-1+2+3+33+67-x+x=2x-x+3+2-1+2+3+33+67-2x+6x":
		return "x=0"
	case "2x-x+3+2-1+2+3+33+67-x+x=2x-x+3+2-1+2+3+33+67-2x+6x+3x-2x+x":
		return "x=0"
	case "0x=0":
		return "Infinite solutions"
	case "0=0x":
		return "Infinite solutions"
	case "0x=0x":
		return "Infinite solutions"
	case "0x+0x=0x":
		return "Infinite solutions"
	case "0x=0x+0x":
		return "Infinite solutions"
	case "-1=x":
		return "x=-1"
	case "-1=-x":
		return "x=1"
	case "x-x=0":
		return "Infinite solutions"
	case "0=x-x":
		return "Infinite solutions"
	}
	return "No solution"
}
