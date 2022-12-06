/*
 * @lc app=leetcode.cn id=809 lang=golang
 *
 * [809] 情感丰富的文字
 */

// @lc code=start
func expressiveWords(s string, words []string) (rs int) {
	for _, word := range words {
		if check(s, word) {
			rs++
		}
	}
	return rs
}

// 检查 一个单词 是否符合s的要求
func check(s, w string) bool {
	lenS := len(s)
	lenW := len(w)

	if lenW > lenS {
		return false
	}

	ks, kw := 0, 0
	for ks < lenS && kw < lenW {
		// fmt.Printf("\n\ns:%c w:%c\n", s[ks], w[kw])
		// 每次判断 字符不一致 直接返回false
		if s[ks] != w[kw] {
			return false
		}
		// 如果相同则计算 在s 当前字符重复的数量
		repeatCntS, lastKeyS := calcRepeatCnt(s, ks)

		// 如果相同则计算 在w 当前字符 重复的数量
		repeatCntW, lastKeyW := calcRepeatCnt(w, kw)

		// "heeelllooo"
		// ["hellllo"]
		// 为避免这种情况  word里重复的cnt > s 重复的数量
		if repeatCntW != repeatCntS && repeatCntS < 3 || repeatCntW > repeatCntS {
			return false
		}
		ks, kw = lastKeyS, lastKeyW
		// fmt.Printf(" lenS:%d lenW:%d lastKeyS:%d lastKeyW:%d \n", lenS, lenW, lastKeyS, lastKeyW)
	}
	return lenS == ks && kw == lenW
}

func calcRepeatCnt(s string, k int) (cnt, lastKey int) {
	lastKey = k
	lenS := len(s)
	for lastKey < lenS && s[k] == s[lastKey] {
		lastKey++
		cnt++
	}
	// fmt.Printf("s:%s c:%c k:%d cnt:%d lastKey:%d \n", s, s[k], k, cnt, lastKey)
	return
}

// @lc code=end

