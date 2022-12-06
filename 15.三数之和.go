/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		// 避免重复的
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		ans = needTwo(nums, i, ans)
	}
	return ans
}

// 剩下的就是2个数 求和 且不重复
func needTwo(nums []int, i int, ans [][]int) [][]int {
	left, right := i+1, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right] + nums[i]

		if sum == 0 {
			fmt.Printf("%d %d %d %d \n", nums[i], nums[left], nums[right], sum)
			ans = append(ans, []int{nums[left], nums[right], nums[i]})
			// 让左侧的指针跳过重复的值。
			tmp := nums[left]
			for tmp == nums[left] && left < right {
				left++
			}
		}

		if sum > 0 {
			right--
		}

		if sum < 0 {
			left++
		}
	}
	return ans
}

// @lc code=end

