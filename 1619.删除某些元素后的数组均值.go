/*
 * @lc app=leetcode.cn id=1619 lang=golang
 *
 * [1619] 删除某些元素后的数组均值
 */

// @lc code=start
func trimMean(arr []int) float64 {

	sort.Ints(arr)
	n := len(arr)

	// 把排序后的 数组掐头去尾 5%
	arr1 := arr[n/20 : 19*n/20]

	sum := 0
	l := len(arr1)
	for i := 0; i < l; i++ {
		sum += arr1[i]
	}

	return float64(sum) / float64(l)
}

// @lc code=end

