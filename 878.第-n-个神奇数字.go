/*
 * @lc app=leetcode.cn id=878 lang=golang
 *
 * [878] 第 N 个神奇数字
 */

// @lc code=start
func nthMagicalNumber(n int, a int, b int) int {
	start := min(a, b)   // 开始的数从两数较小的数开始
	end := n * min(a, b) // 结束的数 到 较小数 * n
	gbs := minGongBeiShu(a, b)

	maxVal := 1000000000 + 7
	cnt := 0

	for start <= end {
		// 找到中间的数,计算出这个数 有几个神奇数字
		mid := (start + end) / 2
		cnt = mid/a + mid/b - mid/gbs
		// 不断调整 左右两边的边界
		if cnt >= n {
			end = mid - 1
		}
		if cnt < n {
			start = mid + 1
		}
		// if cnt == n {
		// 	break
		// }
	}
	return (end + 1) % maxVal

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxGongYueShu(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func minGongBeiShu(a, b int) int {
	return a / maxGongYueShu(a, b) * b
}

// @lc code=end

