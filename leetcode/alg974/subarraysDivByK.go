package main

import "fmt"

/* 给定一个整数数组 A，返回其中元素之和可被 K 整除的（连续、非空）子数组的数目。



示例：

输入：A = [4,5,0,-2,-3,1], K = 5
输出：7
解释：
有 7 个子数组满足其元素之和可被 K = 5 整除：
[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subarray-sums-divisible-by-k
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

// 和可被 K 整除的子数组
func subarraysDivByK(A []int, K int) int {
	// 前缀和，计算所有的前缀和取余结果，若果存在两个相同的取余结果，证明有两个符合条件的结果
	// 前缀和取余结果相同，相减的话就会得到余数0
	res, sum, record := 0, 0, map[int]int{0: 1}
	for i := 0; i < len(A); i++ {
		sum += A[i]
		remainder := (sum%K + K) % K
		res += record[remainder]
		record[remainder]++
	}
	return res
}

func main() {
	res := subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5)
	fmt.Println(res)
}
