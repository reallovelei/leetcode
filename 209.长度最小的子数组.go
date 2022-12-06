/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	// 初始化双指针，从第一个元素开始
	left, right := 0, 0
	sum := 0
	min := math.MaxInt
	for ; right < len(nums); right++ {
		sum += nums[right]
		// 这里可能后面加进来的数会比较大，左侧可以多移动几次。
		// 不能一个简单的 if 判断，所以需要一个循环

		// 当某个数自己就 >= target的时候 要允许left 追上right
		for sum >= target && left <= right {
			fmt.Printf("left:%d right:%d sum:%d min:%d \n", left, right, sum, min)
			tmp := right - left + 1
			if tmp < min {
				min = tmp
			}

			sum -= nums[left] // 将左侧窗口数据移除
			left++            // 左侧窗口向右移动一格
		}
	}
	// 当所有数加起来也 < target的时候 置为0
	if min == math.MaxInt {
		return 0
	}
	return min
}

// @lc code=end

