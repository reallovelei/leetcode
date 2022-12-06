/*
 * @lc app=leetcode.cn id=1652 lang=golang
 *
 * [1652] 拆炸弹
 */

// @lc code=start
func decrypt(code []int, k int) []int {
	rs := make([]int, len(code))
	n := len(code)
	if k == 0 {
		return rs
	}

	code = append(code, code...)
	l, r := 1, k
	if k < 0 {
		// 因为k 是负数，所以left直接 +k
		l, r = n+k, n-1
	}
	fmt.Printf("first l:%d r:%d  \n", l, r)
	fmt.Println(code[l : r+1])
	// 得到第一个数 应该变成 的sum
	sum := 0
	// slice 这种用法 包括left 不包括right
	for _, v := range code[l : r+1] {
		sum += v
	}
	fmt.Println(code, sum)

	for i := 0; i < len(rs); i++ {
		fmt.Printf("l:%d r:%d sum:%d \n", l, r, sum)
		rs[i] = sum
		sum += code[r+1]
		sum -= code[l]
		l, r = l+1, r+1
	}
	return rs
}

// @lc code=end

