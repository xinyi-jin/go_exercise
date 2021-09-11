package alg1720

/* 未知 整数数组 arr 由 n 个非负整数组成。

经编码后变为长度为 n - 1 的另一个整数数组 encoded ，其中 encoded[i] = arr[i] XOR arr[i + 1] 。例如，arr = [1,0,2,1] 经编码后得到 encoded = [1,2,3] 。

给你编码后的数组 encoded 和原数组 arr 的第一个元素 first（arr[0]）。

请解码返回原数组 arr 。可以证明答案存在并且是唯一的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decode-xored-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

// 思路：A^B = C 可以推出 B^C = A 和 A^C = B
func Decode(encoded []int, first int) []int {
	var res = make([]int, len(encoded)+1)
	res[0] = first
	temp := res[0]
	for i := 0; i < len(encoded); i++ {
		res[i+1] = encoded[i] ^ temp
		temp = res[i+1]
	}
	return res
}

// DecodeEx 优化内存
func DecodeEx(encoded []int, first int) []int {
	var res = make([]int, len(encoded)+1)
	res[0] = first
	for i := 0; i < len(encoded); i++ {
		res[i+1] = encoded[i] ^ res[i]
	}
	return res
}
