/*
 * @lc app=leetcode.cn id=1636 lang=golang
 *
 * [1636] 按照频率将数组升序排序
 */

// @lc code=start
func frequencySort(nums []int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	// 将原来的slice 按次数进行排序，这样不用再额外申请存储空间。
	sort.Slice(nums, func(i, j int) bool {
		// 1.这种比较骚操作，不太好理解
		// return m[nums[i]] < m[nums[j]] || (m[nums[i]] == m[nums[j]] && nums[i] > nums[j])

		// 2.这种比较好理解  次数小的放前面
		if m[nums[i]] < m[nums[j]] {
			return true
		}
		// 次数一样， 数大的放前面。
		if m[nums[i]] == m[nums[j]] && nums[i] > nums[j] {
			return true
		}
		return false

	})

	// fmt.Println(nums)
	return nums
}

// @lc code=end

