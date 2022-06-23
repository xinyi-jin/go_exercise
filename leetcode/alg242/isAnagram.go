package alg242

import "sort"

/* 242. 有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。



示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false


提示:

1 <= s.length, t.length <= 5 * 104
s 和 t 仅包含小写字母


进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？ */

// 思路：排序比较
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s1, s2 := SortByByte(s), SortByByte(t)
	sort.Sort(s1)
	sort.Sort(s2)
	return string(s1) == string(s2)
}

type SortByByte []byte

func (a SortByByte) Len() int           { return len(a) }
func (a SortByByte) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByByte) Less(i, j int) bool { return a[i] < a[j] }

func isAnagramEx(s string, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

// 思路：哈希
func isAnagramExtion(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hash := map[rune]int{}
	for _, v := range s {
		hash[v]++
	}
	for _, v := range t {
		if n, exist := hash[v]; exist && n > 0 {
			hash[v]--
			if hash[v] == 0 {
				delete(hash, v)
			}
		} else {
			return false
		}
	}
	return len(hash) == 0
}

// 思路：数组
func isAnagramExtionEx(s string, t string) bool {
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
