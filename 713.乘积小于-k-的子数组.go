/*
 * @lc app=leetcode.cn id=713 lang=golang
 *
 * [713] 乘积小于 K 的子数组
 */

// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) int {
	left, right, ans, product := 0, 0, 0, 1
	n := len(nums)
	for ; right < n; right++ {
		product = product * nums[right]

		for product >= k && left <= right {
			product = product / nums[left]
			left++
		}
		// 右侧不变  左侧变，几个数就有几个
		ans += right - left + 1
	}
	return ans
}

// @lc code=end

