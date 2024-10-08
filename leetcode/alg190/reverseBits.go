package leetcode

/* 颠倒给定的 32 位无符号整数的二进制位。



提示：

请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在上面的 示例 2 中，输入表示有符号整数 -3，输出表示有符号整数 -1073741825。


进阶:
如果多次调用这个函数，你将如何优化你的算法？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-bits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

// 思路：使用位运算(4ms、2.6mb)
func ReverseBits(num uint32) uint32 {
	for i := 0; i < 16; i++ {
		if num>>i&1 != num>>(32-i-1)&1 {
			num ^= 1 << i
			num ^= 1 << (32 - i - 1)
		}
	}
	return num
}

// 使用栈(执行效率明显高于单纯位运算，0ms、2.6mb)
func ReverseBits1(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res |= (num >> i & 1) << (32 - i - 1)
	}
	return res
}
