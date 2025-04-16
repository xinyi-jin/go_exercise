package alg2179

import (
	"fmt"
)

/* 2179. 统计数组中好三元组数目
困难
相关标签
相关企业
提示
给你两个下标从 0 开始且长度为 n 的整数数组 nums1 和 nums2 ，两者都是 [0, 1, ..., n - 1] 的 排列 。

好三元组 指的是 3 个 互不相同 的值，且它们在数组 nums1 和 nums2 中出现顺序保持一致。换句话说，如果我们将 pos1v 记为值 v 在 nums1 中出现的位置，pos2v 为值 v 在 nums2 中的位置，那么一个好三元组定义为 0 <= x, y, z <= n - 1 ，且 pos1x < pos1y < pos1z 和 pos2x < pos2y < pos2z 都成立的 (x, y, z) 。

请你返回好三元组的 总数目 。



示例 1：

输入：nums1 = [2,0,1,3], nums2 = [0,1,2,3]
输出：1
解释：
总共有 4 个三元组 (x,y,z) 满足 pos1x < pos1y < pos1z ，分别是 (2,0,1) ，(2,0,3) ，(2,1,3) 和 (0,1,3) 。
这些三元组中，只有 (0,1,3) 满足 pos2x < pos2y < pos2z 。所以只有 1 个好三元组。
示例 2：

输入：nums1 = [4,0,1,3,2], nums2 = [4,1,0,2,3]
输出：4
解释：总共有 4 个好三元组 (4,0,3) ，(4,0,2) ，(4,1,3) 和 (4,1,2) 。


提示：

n == nums1.length == nums2.length
3 <= n <= 105
0 <= nums1[i], nums2[i] <= n - 1
nums1 和 nums2 是 [0, 1, ..., n - 1] 的排列。 */

// 思路：同简单题一样 枚举 遍历尝试
// 超时 因为是困难等级的，尝试其他思路
func goodTripletsE(nums1 []int, nums2 []int) int64 {
	n := int64(0)
	l := len(nums1)
	for i := 0; i <= l-3; i++ {
		for j := i; j <= l-2; j++ {
			if nums1[i] != nums1[j] {
				for k := j; k <= l-1; k++ {
					if nums1[j] != nums1[k] && hasGoodTriplets(nums1[i], nums1[j], nums1[k], nums2) {
						n++
					}
				}
			}
		}
	}
	return n
}

func hasGoodTriplets(x, y, z int, nums []int) bool {
	n := 0
	for i := 0; i < len(nums); i++ {
		switch n {
		case 0:
			if nums[i] == x {
				n++
			}
		case 1:
			if nums[i] == y {
				n++
			}
		case 2:
			if nums[i] == z {
				n++
			}
		}
	}
	return n == 3
}

// 思路：利用轮廓画出好三元数组的形状
// 1. 找到好三元数组
// 2. 好三元数组转换字符串拼接形成画像
// 3. 比对画像数据 得到个数
func goodTripletsEx(nums1 []int, nums2 []int) int64 {
	n := int64(0)
	l := len(nums1)
	hxMap := make(map[string]int)
	for i := 0; i <= l-3; i++ {
		for j := i; j <= l-2; j++ {
			if nums1[i] != nums1[j] {
				for k := j; k <= l-1; k++ {
					if nums1[j] != nums1[k] {
						hxStr := fmt.Sprintf("%d-%d-%d", nums1[i], nums1[j], nums1[k])
						hxMap[hxStr]++
					}
				}
			}
		}
	}
	for i := 0; i <= l-3; i++ {
		for j := i; j <= l-2; j++ {
			if nums2[i] != nums2[j] {
				for k := j; k <= l-1; k++ {
					if nums2[j] != nums2[k] {
						hxStr := fmt.Sprintf("%d-%d-%d", nums2[i], nums2[j], nums2[k])
						if hxMap[hxStr] > 0 {
							n++
						}
					}
				}
			}
		}
	}
	return n
}

// 官方题解：树状数组
type FenwickTree struct {
	tree []int
}

func fenwickTree(size int) *FenwickTree {
	return &FenwickTree{tree: make([]int, size+1)}
}

func (ft *FenwickTree) update(index, delta int) {
	index++
	for index < len(ft.tree) {
		ft.tree[index] += delta
		index += index & -index
	}
}

func (ft *FenwickTree) query(index int) int {
	index++
	res := 0
	for index > 0 {
		res += ft.tree[index]
		index -= index & -index
	}
	return res
}

func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	pos2, reversedIndexMapping := make([]int, n), make([]int, n)
	for i, num := range nums2 {
		pos2[num] = i
	}

	for i, num := range nums1 {
		reversedIndexMapping[pos2[num]] = i
	}
	tree := fenwickTree(n)
	var res int64
	for value := 0; value < n; value++ {
		pos := reversedIndexMapping[value]
		left := tree.query(pos)
		tree.update(pos, 1)
		right := (n - 1 - pos) - (value - left)
		res += int64(left * right)
	}
	return res
}
