/*
 * @lc app=leetcode.cn id=1710 lang=golang
 *
 * [1710] 卡车上的最大单元数
 */

// @lc code=start
func maximumUnits(boxTypes [][]int, truckSize int) int {
	rs := 0
	sort.Slice(boxTypes, func(a, b int) bool {
		return boxTypes[a][1] > boxTypes[b][1]
	})

	for _, val := range boxTypes {
		if 0 >= truckSize {
			break
		}

		// 能把这一组的箱子 都装完。
		if val[0] < truckSize {
			rs = rs + val[0]*val[1]
			truckSize = truckSize - val[0]
		} else { // 位置不够了，只能装一部分箱子了，不能全部都上车。
			rs = rs + truckSize*val[1]
			truckSize = 0
		}
	}

	return rs
}

// @lc code=end

