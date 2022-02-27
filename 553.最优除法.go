/*
 * @lc app=leetcode.cn id=553 lang=golang
 *
 * [553] 最优除法
求解思路
最大的结果 = 分子越大 / 分母越小
分析可以知道nums[0]一定是分子不可能变化，nums[1]一定为分母，需要找到如何使分母最大
所以最终只要把第一个元素后的所有元素括起来就能够保证最终的值是最大的

*/

// @lc code=start
func optimalDivision(nums []int) string {
	cnt := len(nums)
	var ans string
	// 如果只有一个元素 直接返回
	if cnt == 1 {
		return strconv.Itoa(nums[0])
	}

	// 如果只有两个元素 直接返回 2个数的相除的字符串
	if cnt == 2 {
		return fmt.Sprintf("%d / %d", nums[0], nums[1])
	}

	// 先把前2个拼好 加上左侧括号
	ans = fmt.Sprintf("%d / ( %d", nums[0], nums[1])

	// 拼装中间的数， 从3个数开始
	for _, v := range nums[2:] {
		ans += fmt.Sprintf("/ %d", v)
	}
	// 追加结尾右侧括号
	ans += fmt.Sprintf(")")
	return ans
}

// @lc code=end

