/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	hash := map[string][]string{}
	ans := [][]string{}

	for _, s := range strs {
		bs := []byte(s)
		sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
		sortStr := string(bs)
		hash[sortStr] = append(hash[sortStr], s)

	}

	for _, v := range hash {
		ans = append(ans, v)
	}
	return ans
}

// @lc code=end

