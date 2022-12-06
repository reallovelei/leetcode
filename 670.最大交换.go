/*
 * @lc app=leetcode.cn id=670 lang=golang
 *
 * [670] 最大交换
 */

// @lc code=start
func maximumSwap(num int) int {
	if num < 11 {
		return num
	}
	rs := num
	// 将每个位上的数字 转成数组
	s := []byte(strconv.Itoa(num))
	n := len(s)

	//默认个位的数最大
	idxMax, idx1, idx2 := n-1, -1, -1
	// 从最右边遍历，
	for i := n - 1; i >= 0; i-- {
		if s[i] > s[idxMax] {
			idxMax = i
		} else if s[i] < s[idxMax] {
			idx1, idx2 = i, idxMax
			// fmt.Printf("idx1:%d, idx2:%d \n", idx1, idx2)
		}
	}

	// fmt.Printf("idx1:%d, idx2:%d idxMax:%d \n", idx1, idx2, idxMax)
	// 如果没有改变直接返回
	if idx1 < 0 {
		return num
	}
	s[idx1], s[idx2] = s[idx2], s[idx1]

	rs, _ = strconv.Atoi(string(s))
	return rs

}

// @lc code=end

