/*
 * @lc app=leetcode.cn id=2032 lang=golang
 *
 * [2032] 至少在两个数组中出现的值
 */

// @lc code=start
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	// 整体思路:
	// 共三个数组，用第三位来表示在哪个数组里
	// 遍历hash  如果 cnt&(cnt-1) ＞0 代表在2个以上的数组里。 这个判断比较巧妙
	cntHash := make(map[int]int, len(nums1))
	ans := []int{}
	for k, nums := range [][]int{nums1, nums2, nums3} {
		for _, val := range nums {
			// 3位 分表代表 在哪个数组里 111
			cntHash[val] |= 1 << k
		}
	}

	for idx, cnt := range cntHash {
		if cnt&(cnt-1) > 0 {
			ans = append(ans, idx)
		}
	}
	return ans
}

// @lc code=end

