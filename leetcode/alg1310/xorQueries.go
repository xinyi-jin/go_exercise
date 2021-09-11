package alg1310

/* 1310. 子数组异或查询
有一个正整数数组 arr，现给你一个对应的查询数组 queries，其中 queries[i] = [Li, Ri]。

对于每个查询 i，请你计算从 Li 到 Ri 的 XOR 值（即 arr[Li] xor arr[Li+1] xor ... xor arr[Ri]）作为本次查询的结果。

并返回一个包含给定查询 queries 所有结果的数组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/xor-queries-of-a-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

// 思路：根据意思直接求解(超时)
func XorQueries(arr []int, queries [][]int) []int {
	l := len(queries)
	if l < 1 {
		return nil
	}
	resArr := make([]int, l)
	for i := 0; i < l; i++ {
		for j := queries[i][0]; j <= queries[i][1]; j++ {
			resArr[i] ^= arr[j]
		}
	}
	return resArr
}

// 优化版：先生成前缀异或数组，所求取件的异或最终值为：前缀异或数组开始下标值 ^ （前缀异或数组结束下标值+1）
func XorQueriesEx(arr []int, queries [][]int) []int {
	xorPre := make([]int, len(arr)+1)
	for i, v := range arr {
		xorPre[i+1] = xorPre[i] ^ v
	}
	resArr := make([]int, len(queries))
	for i, v := range queries {
		resArr[i] = xorPre[v[0]] ^ xorPre[v[1]+1]
	}
	return resArr
}
