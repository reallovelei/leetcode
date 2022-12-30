/*
 * @lc app=leetcode.cn id=2037 lang=golang
 *
 * [2037] 使每位学生都有座位的最少移动次数
 */

// @lc code=start
func minMovesToSeat(seats []int, students []int) int {
	ans := 0
	sort.Ints(seats)
	sort.Ints(students)
	for k, v := range seats {
		ans += abs(v - students[k])
	}
	// fmt.Println(seats, students)
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

