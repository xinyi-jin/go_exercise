package main

/* 给定一个正整数数组 A，如果 A 的某个子数组中不同整数的个数恰好为 K，则称 A 的这个连续、不一定独立的子数组为好子数组。

（例如，[1,2,3,1,2] 中有 3 个不同的整数：1，2，以及 3。）

返回 A 中好子数组的数目。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subarrays-with-k-different-integers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

// 未完善代码
/* func subarraysWithKDistinct(A []int, K int) int {
	n := len(A)
	if n < 2 {
		return 1
	}
	left, right := 0, 1
	res := 0
	cnt := 0
	index := 0
	for right < n {
		if A[right] > A[right-1] {
			cnt++
		} else {
			index = right
			flag := false
			for _, v := range A[:index] {
				if A[right] == v {
					flag = true
				}
			}
			if flag == false {
				cnt++
			}
		}
		if cnt == K {
			res++
		}
		if cnt > K {
			left++
			right = left
		}
		right++
	}
	return res
} */

func subarraysWithKDistinct(A []int, K int) (ans int) {
	n := len(A)
	num1 := make([]int, n+1)
	num2 := make([]int, n+1)
	var tot1, tot2, left1, left2 int
	for _, v := range A {
		if num1[v] == 0 {
			tot1++
		}
		num1[v]++
		if num2[v] == 0 {
			tot2++
		}
		num2[v]++
		for tot1 > K {
			num1[A[left1]]--
			if num1[A[left1]] == 0 {
				tot1--
			}
			left1++
		}
		for tot2 > K-1 {
			num2[A[left2]]--
			if num2[A[left2]] == 0 {
				tot2--
			}
			left2++
		}
		ans += left2 - left1
	}
	return ans
}

func main() {

}
