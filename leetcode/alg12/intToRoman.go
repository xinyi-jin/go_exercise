package alg12

/* 12. 整数转罗马数字
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

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
给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-to-roman
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

// 思路：根据每个位上的值逐个翻译 (未完成，存在bug)
func IntToRoman(num int) string {
	res := ""
	for i := 1000; i >= 1; i /= 10 {
		bn := num / i * i
		if bn > 1 {
			AddBitValue(&res, bn)
		}
	}
	return res
}
func AddBitValue(str *string, num int) {
	for i := 0; i < num; i++ {
		if num < 1 {
			return
		}
		switch {
		case num >= 1000:
			*str += "M"
			num -= 1000
		case num >= 900:
			*str += "CM"
			num -= 900
		case num >= 500:
			*str += "CM"
			num -= 500
		case num >= 400:
			*str += "CD"
			num -= 400
		case num >= 100:
			*str += "C"
			num -= 100
		case num >= 90:
			*str += "XC"
			num -= 90
		case num >= 50:
			*str += "L"
			num -= 50
		case num >= 40:
			*str += "XL"
			num -= 40
		case num >= 10:
			*str += "X"
			num -= 10
		case num >= 9:
			*str += "IX"
			num -= 9
		case num >= 5:
			*str += "V"
			num -= 5
		case num >= 4:
			*str += "IV"
			num -= 4
		case num >= 1:
			*str += "I"
			num -= 1
		}
	}
}

var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// 通过减法计算得出每位的值
func intToRoman(num int) string {
	roman := []byte{}
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			roman = append(roman, vs.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}

var (
	thousands = []string{"", "M", "MM", "MMM"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

// 采用硬编码方式（列举所有可能）
func IntToRomanEx(num int) string {
	return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}
