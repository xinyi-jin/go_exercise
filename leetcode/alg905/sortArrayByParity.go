package alg905

/* 905. 按奇偶排序数组
给你一个整数数组 nums，将 nums 中的的所有偶数元素移动到数组的前面，后跟所有奇数元素。

返回满足此条件的 任一数组 作为答案。



示例 1：

输入：nums = [3,1,2,4]
输出：[2,4,3,1]
解释：[4,2,3,1]、[2,4,1,3] 和 [4,2,1,3] 也会被视作正确答案。
示例 2：

输入：nums = [0]
输出：[0]


提示：

1 <= nums.length <= 5000
0 <= nums[i] <= 5000 */
// 思路：快慢指针
func sortArrayByParity(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	for i := 0; i < n; i++ {
		if nums[i]&1 == 0 {
			continue
		}
		for j := i + 1; j < n; j++ {
			if nums[j]&1 == 0 {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
	return nums
}

// 双指针
func sortArrayByParityEx(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	left, right := 0, n-1
	for _, v := range nums {
		if v&1 == 0 {
			res[left] = v
			left++
		} else {
			res[right] = v
			right--
		}
	}
	return res
}
