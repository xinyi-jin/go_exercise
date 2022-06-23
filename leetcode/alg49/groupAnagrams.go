package alg49

/* 49. 字母异位词分组
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。



示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]


提示：

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母 */

// 思路：滑动窗口
func groupAnagrams(strs []string) (res [][]string) {
	for i := 0; i < len(strs); i++ {
		temp := []string{strs[i]}
		for j := i + 1; j < len(strs); j++ {
			if isAnagram(strs[i], strs[j]) {
				temp = append(temp, strs[j])
				strs = append(strs[:j], strs[j+1:]...)
				j--
			}
		}
		res = append(res, temp)
	}
	return
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	array := [26]int{}
	for _, v := range s {
		array[v-'a']++
	}
	for _, v := range t {
		array[v-'a']--
		if array[v-'a'] < 0 {
			return false
		}
	}
	return true
}
