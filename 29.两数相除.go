/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(a int, b int) int {
	rs := 0
	sign := false

	if (a < 0 && b > 0) || (a > 0 && b < 0) {
		sign = true
	}

	a, b = abs(a), abs(b)
	if a < b {
		return 0
	}

	for i := 31; i >= 0; i-- {
		// 正数 溢出处理
		if i == 31 && sign == true {
			return 1<<i - 1
		}

		a -= b << i
		rs += 1 << i
	}

	if rs >= 1<<31 && sign == false {
		rs = 1 << 31
	}

	if !sign {
		return -rs
	}
	return rs
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end

