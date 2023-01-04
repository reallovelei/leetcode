/*
 * @lc app=leetcode.cn id=1802 lang=golang
 *
 * [1802] 有界数组中指定下标处的最大值
 */

// @lc code=start
func maxValue(n int, index int, maxSum int) int {
	sum := 0

	// 这里是二分查找
	return sort.Search(maxSum, func(i int) bool {
		i++
		// 左边的数字和 + 右边的数字和
		left := subSum(i-1, index)
		right := subSum(i, n-index)
		sum = left + right
		return sum > maxSum
	})

	// 线性查找  timelimit了
	// for sum <= maxSum {
	// 	i++
	// 	left := subSum(i-1, index)
	// 	right := subSum(i, n-index)
	// 	sum = left + right
	// }
	// fmt.Println(i, sum)
	// return i - 1
}

func subSum(maxValue int, cnt int) int {
	// 节点数量 大于 最大的元素值
	if cnt > maxValue {
		// 会多出来 (cnt - maxValue) 个1
		// 首元素为 肯定为1 加 尾元素就是最大值      剩余这么多个1
		return (1+maxValue)*maxValue/2 + (cnt - maxValue)
	} else {
		// 543 这种情况 少 maxValue - cnt个 1
		// 首元素为 maxValue - cnt + 1, 尾元素为maxValue
		return (maxValue - cnt + 1 + maxValue) * cnt / 2
	}
}

// @lc code=end

