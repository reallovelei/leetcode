/*
 * @lc app=leetcode.cn id=1779 lang=golang
 *
 * [1779] 找到最近的有相同 X 或 Y 坐标的点
 */

// @lc code=start
func nearestValidPoint(x int, y int, points [][]int) int {
	min, mhd := math.MaxInt, 0
	minKey := -1
	for k, point := range points {
		// 判断为有效点
		if point[0] == x || point[1] == y {
			mhd = abs(x-point[0]) + abs(y-point[1])
			if mhd < min {
				min = mhd
				minKey = k
			}
		}
	}
	return minKey
}

// 绝对值
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// @lc code=end

