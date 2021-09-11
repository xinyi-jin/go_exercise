package alg1486

/* 1486. 数组异或操作
给你两个整数，n 和 start 。

数组 nums 定义为：nums[i] = start + 2*i（下标从 0 开始）且 n == nums.length 。

请返回 nums 中所有元素按位异或（XOR）后得到的结果。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/xor-operation-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

// 思路：直接计算
func XorOperation(n int, start int) int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = start + 2*i
	}

	res := arr[0]
	for i := 1; i < len(arr); i++ {
		res ^= arr[i]
	}
	return res
}

// 优化内存 (时间复杂度 O(n) 空间复杂度 O(1))
func XorOperationEx(n int, start int) int {
	res := start
	for i := 1; i < n; i++ {
		res ^= (start + 2*i)
	}
	return res
}

// 进一步优化(其实这样做的效果和上边的实现没啥差别)
func XorOperationExtion(n int, start int) int {
	res := start
	for i := 1; i < n; i++ {
		start += 2
		res ^= start
	}
	return res
}

func sumXor(x int) int {
	switch x % 4 {
	case 0:
		return x
	case 1:
		return 1
	case 2:
		return x + 1
	default:
		return 0
	}
}

// 官方题解，时间复杂度和空间复杂度均是O(1) (效率和优化版本代码基本一致，应该是系数n没有达到差异的数值)
func XorOperationGF(n, start int) (ans int) {
	s, e := start>>1, n&start&1
	ret := sumXor(s-1) ^ sumXor(s+n-1)
	return ret<<1 | e
}
