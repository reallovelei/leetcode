/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	ak, bk := len(a)-1, len(b)-1
	ans := ""
	carry := 0

	for ak >= 0 || bk >= 0 {
		av, bv := 0, 0
		// carry = 0
		if ak >= 0 {
			av = int(a[ak] - '0')
		}
		if bk >= 0 {
			bv = int(b[bk] - '0')
		}

		tmp := av + bv + carry
		// fmt.Printf("av:%d bv:%d tmp:%d \n", av, bv, tmp)
		carry = tmp / 2
		tmp = tmp % 2
		ans = strconv.Itoa(tmp) + ans

		ak--
		bk--
	}

	if carry == 1 {
		ans = "1" + ans
	}
	return ans
}

// @lc code=end

