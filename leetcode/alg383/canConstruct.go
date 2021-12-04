package alg383

/* 383. 赎金信
为了不在赎金信中暴露字迹，从杂志上搜索各个需要的字母，组成单词来表达意思。

给你一个赎金信 (ransomNote) 字符串和一个杂志(magazine)字符串，判断 ransomNote 能不能由 magazines 里面的字符构成。

如果可以构成，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。



示例 1：

输入：ransomNote = "a", magazine = "b"
输出：false
示例 2：

输入：ransomNote = "aa", magazine = "ab"
输出：false
示例 3：

输入：ransomNote = "aa", magazine = "aab"
输出：true


提示：

1 <= ransomNote.length, magazine.length <= 105
ransomNote 和 magazine 由小写英文字母组成 */

// 思路：直接使用map计算两个所含各个小写字母的个数
func canConstruct(ransomNote string, magazine string) bool {
	magazineMap, ransomNoteMap := make(map[rune]int), make(map[rune]int)
	for _, v := range magazine {
		magazineMap[v]++
	}
	for _, v := range ransomNote {
		ransomNoteMap[v]++
	}
	for k, v := range ransomNoteMap {
		if magazineMap[k] < v {
			return false
		}
	}
	return true
}

// 其他人提交解法（优化内存）
func canConstructEx(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	var cnt [26]int
	for _, v := range magazine {
		cnt[v-'a']++
	}
	for _, v := range ransomNote {
		cnt[v-'a']--
		if cnt[v-'a'] < 0 {
			return false
		}
	}
	return true
}
