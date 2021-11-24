package alg423

import "bytes"

/* 423. 从英文中重建数字
给你一个字符串 s ，其中包含字母顺序打乱的用英文单词表示的若干数字（0-9）。按 升序 返回原始的数字。



示例 1：

输入：s = "owoztneoer"
输出："012"
示例 2：

输入：s = "fviefuro"
输出："45"


提示：

1 <= s.length <= 105
s[i] 为 ["e","g","f","i","h","o","n","s","r","u","t","w","v","x","z"] 这些字符之一
s 保证是一个符合题目要求的字符串 */

// 官方题解：脑筋急转弯。
/* 方法一：依次确定每一个数字的次数
思路与算法

首先我们可以统计每个字母分别在哪些数字中出现：

字母	数字
e	0 1 3 5 7 8 9
f	4 5
g	8
h	3 8
i	5 6 8 9
n	1 7 9
o	0 1 2 4
r	0 3 4
s	6 7
t	2 3 8
u	4
v	5 7
w	2
x	6
z	0
可以发现，\text{z, w, u, x, g}z, w, u, x, g 都只在一个数字中，即 0, 2, 4, 6, 80,2,4,6,8 中出现。因此我们可以使用一个哈希表统计每个字母出现的次数，那么 \text{z, w, u, x, g}z, w, u, x, g 出现的次数，即分别为 0, 2, 4, 6, 80,2,4,6,8 出现的次数。

随后我们可以注意那些只在两个数字中出现的字符：

\text{h}h 只在 3, 83,8 中出现。由于我们已经知道了 88 出现的次数，因此可以计算出 33 出现的次数。

\text{f}f 只在 4, 54,5 中出现。由于我们已经知道了 44 出现的次数，因此可以计算出 55 出现的次数。

\text{s}s 只在 6, 76,7 中出现。由于我们已经知道了 66 出现的次数，因此可以计算出 77 出现的次数。

此时，我们还剩下 11 和 99 的出现次数没有求出：

\text{o}o 只在 0, 1, 2, 40,1,2,4 中出现，由于我们已经知道了 0, 2, 40,2,4 出现的次数，因此可以计算出 11 出现的次数。
最后的 99 就可以通过 \text{n, i, e}n, i, e 中的任一字符计算得到了。这里推荐使用 \text{i}i 进行计算，因为 \text{n}n 在 99 中出现了 22 次，\text{e}e 在 33 中出现了 22 次，容易在计算中遗漏。

当我们统计完每个数字出现的次数后，我们按照升序将它们进行拼接即可。

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/reconstruct-original-digits-from-english/solution/cong-ying-wen-zhong-zhong-jian-shu-zi-by-9g1r/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。 */
func originalDigits(s string) string {
	cnt := make([]int, 10)
	res := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		res[s[i]]++
	}
	cnt[0] = res['z']
	cnt[2] = res['w']
	cnt[4] = res['u']
	cnt[6] = res['x']
	cnt[8] = res['g']

	cnt[3] = res['h'] - cnt[8]
	cnt[5] = res['f'] - cnt[4]
	cnt[7] = res['s'] - cnt[6]

	cnt[1] = res['o'] - cnt[0] - cnt[2] - cnt[4]
	cnt[9] = res['i'] - cnt[5] - cnt[6] - cnt[8]

	ans := make([]byte, 0)
	for i, n := range cnt {
		ans = append(ans, bytes.Repeat([]byte{byte('0' + i)}, n)...)
	}
	return string(ans)
}

func getNumber(s string) map[byte]int {
	res := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		res[s[i]]++
	}
	return res
}
