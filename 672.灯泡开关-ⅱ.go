/*
 * @lc app=leetcode.cn id=672 lang=golang
 *
 * [672] 灯泡开关 Ⅱ
 */

// @lc code=start
func flipLights(n int, presses int) int {
	if presses == 0 {
		return 1
	}
	// 只有一盏灯
	if n == 1 { // 0 | 1
		return 2
	}

	if n == 2 {
		if presses == 1 { // 00/01/10
			return 3
		} else { // 00/01/10/11
			return 4
		}
	}
	// 灯的数量 >= 3
	if presses == 1 {
		return 4
	}
	if presses == 2 {
		return 7
	} else {
		return 8
	}

}

// @lc code=end

