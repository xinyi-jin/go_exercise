package alg504

import (
	"fmt"
	"math"
	"strconv"
)

/* 504. 七进制数
给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。



示例 1:

输入: num = 100
输出: "202"
示例 2:

输入: num = -7
输出: "-10"


提示：

-107 <= num <= 107 */

// 思路:根据进制转换原理，直接计算
func convertToBase7(num int) string {
	s, zf, res := "", "", ""
	if num < 0 {
		zf += "-"
	}
	for {
		s += fmt.Sprint(math.Abs(float64(num % 7)))
		num = num / 7
		if num == 0 {
			s += zf
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}

// 直接调用API
func convertToBase7Ex(num int) string {
	return strconv.FormatInt(int64(num), 7)
}
