/*
 * @lc app=leetcode.cn id=1742 lang=golang
 *
 * [1742] 盒子中小球的最大数量
 */

// @lc code=start
func countBalls(lowLimit int, highLimit int) int {
	cntMap := make(map[int]int, 10)
	ans := 0
	for i := lowLimit; i <= highLimit; i++ {
		num := 0
		for j := i; j > 0; j = j / 10 {
			num = num + j%10
		}
		cntMap[num]++
		// fmt.Println(num, cntMap)
		ans = max(ans, cntMap[num])
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

