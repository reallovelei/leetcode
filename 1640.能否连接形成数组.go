/*
 * @lc app=leetcode.cn id=1640 lang=golang
 *
 * [1640] 能否连接形成数组
 */

// @lc code=start
func canFormArray(arr []int, pieces [][]int) bool {
	// 先遍历 pieces, 以每个数组的首元素为key
	mp := make(map[int]int, len(pieces))
	for kp, vp := range pieces {
		mp[vp[0]] = kp
	}

	for ka := 0; ka < len(arr); {
		// 先拿到 arr 里的元素，再到pieces里找 对应的首元素为key的数组。
		va, ok := arr[ka]
		if !ok { // 避免从pieces里遍历 出来，在原数组找不到对应的元素了
			return false
		}

		vp, ok := pieces[va]
		if !ok {
			return false
		}

		for k, v := range vp {
			if v != arr[ka] {
				return false
			}
			ka++
		}
		return true
	}
}

// @lc code=end

