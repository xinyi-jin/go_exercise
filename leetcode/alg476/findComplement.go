package alg476

/* 476. 数字的补数
对整数的二进制表示取反（0 变 1 ，1 变 0）后，再转换为十进制表示，可以得到这个整数的补数。

例如，整数 5 的二进制表示是 "101" ，取反后得到 "010" ，再转回十进制表示得到补数 2 。
给你一个整数 num ，输出它的补数。



示例 1：

输入：num = 5
输出：2
解释：5 的二进制表示为 101（没有前导零位），其补数为 010。所以你需要输出 2 。
示例 2：

输入：num = 1
输出：0
解释：1 的二进制表示为 1（没有前导零位），其补数为 0。所以你需要输出 0 。


提示：

1 <= num < 231 */

/*
解题思路：
末位向高位逐个计算
最后把所有位数上都置值为1 与原数字 进行异或操作
*/
func findComplement(num int) int {
	temp, c := num, 0
	for temp > 0 {
		temp >>= 1
		c = (c << 1) + 1
	}
	return num ^ c
}

// 性能方面与上方写法基本没区别
func findComplementEx(num int) int {
	temp, c := num, 0
	for temp > 0 {
		temp >>= 1
		c++
	}
	return num ^ (1<<c - 1)
}

// 性能方面与上方写法基本没区别
func findComplementExt(num int) int {
	if num < 2 {
		return 1 - num
	} else if num < 4 {
		return 3 - num
	} else if num < 8 {
		return 7 - num
	} else if num < 16 {
		return 15 - num
	} else if num < 32 {
		return 31 - num
	} else if num < 64 {
		return 63 - num
	} else if num < 128 {
		return 127 - num
	} else if num < 256 {
		return 255 - num
	} else if num < 512 {
		return 511 - num
	} else if num < 1024 {
		return 1023 - num
	} else if num < 2048 {
		return 2047 - num
	} else if num < 4096 {
		return 4095 - num
	} else if num < 8192 {
		return 8191 - num
	} else if num < 16384 {
		return 16383 - num
	} else if num < 32768 {
		return 32767 - num
	} else if num < 65536 {
		return 65535 - num
	} else if num < 131072 {
		return 131071 - num
	} else if num < 262144 {
		return 262143 - num
	} else if num < 524288 {
		return 524287 - num
	} else if num < 1048576 {
		return 1048575 - num
	} else if num < 2097152 {
		return 2097151 - num
	} else if num < 4194304 {
		return 4194303 - num
	} else if num < 8388608 {
		return 8388607 - num
	} else if num < 16777216 {
		return 16777215 - num
	} else if num < 33554432 {
		return 33554431 - num
	} else if num < 67108864 {
		return 67108863 - num
	} else if num < 134217728 {
		return 134217727 - num
	} else if num < 268435456 {
		return 268435455 - num
	} else if num < 536870912 {
		return 536870911 - num
	} else if num < 1073741824 {
		return 1073741823 - num
	} else {
		return 2147483647 - num
	}
}

// 性能方面与上方写法基本没区别
func findComplementExti(num int) int {
	temp := num
	temp |= temp >> 1
	temp |= temp >> 2
	temp |= temp >> 4
	temp |= temp >> 8
	temp |= temp >> 16
	return num ^ temp
}
