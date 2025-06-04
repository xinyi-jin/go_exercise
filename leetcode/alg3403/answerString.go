package alg3403

/* 3403. 从盒子中找出字典序最大的字符串 I
中等
相关标签
premium lock icon
相关企业
提示
给你一个字符串 word 和一个整数 numFriends。

Alice 正在为她的 numFriends 位朋友组织一个游戏。游戏分为多个回合，在每一回合中：

word 被分割成 numFriends 个 非空 字符串，且该分割方式与之前的任意回合所采用的都 不完全相同 。
所有分割出的字符串都会被放入一个盒子中。
在所有回合结束后，找出盒子中 字典序最大的 字符串。



示例 1：

输入: word = "dbca", numFriends = 2

输出: "dbc"

解释:

所有可能的分割方式为：

"d" 和 "bca"。
"db" 和 "ca"。
"dbc" 和 "a"。
示例 2：

输入: word = "gggg", numFriends = 4

输出: "g"

解释:

唯一可能的分割方式为："g", "g", "g", 和 "g"。



提示:

1 <= word.length <= 5 * 103
word 仅由小写英文字母组成。
1 <= numFriends <= word.length */

// 思路：直接查找最长子串，判断长度
func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	res := ""
	n := len(word)
	for i := 0; i < n; i++ {
		end := i + n - numFriends + 1
		if end > n {
			end = n
		}
		str := word[i:end]
		if str > res {
			res = str
		}
	}
	return res
}
