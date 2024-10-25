package alg1657

import "sort"

/* 1657. 确定两个字符串是否接近
提示
中等
90
相关企业
如果可以使用以下操作从一个字符串得到另一个字符串，则认为两个字符串 接近 ：

操作 1：交换任意两个 现有 字符。
例如，abcde -> aecdb
操作 2：将一个 现有 字符的每次出现转换为另一个 现有 字符，并对另一个字符执行相同的操作。
例如，aacabb -> bbcbaa（所有 a 转化为 b ，而所有的 b 转换为 a ）
你可以根据需要对任意一个字符串多次使用这两种操作。

给你两个字符串，word1 和 word2 。如果 word1 和 word2 接近 ，就返回 true ；否则，返回 false 。



示例 1：

输入：word1 = "abc", word2 = "bca"
输出：true
解释：2 次操作从 word1 获得 word2 。
执行操作 1："abc" -> "acb"
执行操作 1："acb" -> "bca"
示例 2：

输入：word1 = "a", word2 = "aa"
输出：false
解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。
示例 3：

输入：word1 = "cabbba", word2 = "abbccc"
输出：true
解释：3 次操作从 word1 获得 word2 。
执行操作 1："cabbba" -> "caabbb"
执行操作 2："caabbb" -> "baaccc"
执行操作 2："baaccc" -> "abbccc"
示例 4：

输入：word1 = "cabbba", word2 = "aabbss"
输出：false
解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。


提示：

1 <= word1.length, word2.length <= 105
word1 和 word2 仅包含小写英文字母 */

// 思路：
// 1. 判断长度是否匹配
// 2. 判断元素字符类型是否匹配
// 3. 判断元素字符类型对应个数是否可相互转换
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	cntMap1, cntMap2 := make(map[rune]int), make(map[rune]int)
	for _, v := range word1 {
		cntMap1[v]++
	}
	for _, v := range word2 {
		cntMap2[v]++
	}
	if len(cntMap1) != len(cntMap2) {
		return false
	}
	for k := range cntMap1 {
		if _, ok := cntMap2[k]; !ok {
			return false
		}
	}

	arr1, arr2 := make([]int, 0), make([]int, 0)
	for _, v := range cntMap1 {
		arr1 = append(arr1, v)
	}
	for _, v := range cntMap2 {
		arr2 = append(arr2, v)
	}
	sort.Ints(arr1)
	sort.Ints(arr2)
	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}
	return true
}
