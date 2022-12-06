/*
 * @lc app=leetcode.cn id=1608 lang=golang
 *
 * [1608] 特殊数组的特征值
 */

// @lc code=start
func specialArray(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Println(nums)

	len := len(nums)
	if nums[0] == 0 {
		return -1
	}

	for i := 1; i <= len; i++ {
		fmt.Printf("i:%d n:%d \n", i, nums[i-1])
		// 前一个数 >= i
		if nums[i-1] >= i && (i == len || nums[i] < i) {
			return i
		}
	}
	return -1
}

// @lc code=end

