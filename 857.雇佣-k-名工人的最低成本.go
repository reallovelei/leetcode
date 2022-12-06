/*
 * @lc app=leetcode.cn id=857 lang=golang
 *
 * [857] 雇佣 K 名工人的最低成本
 */

type mHeap struct {
	q []int
}

func (h *mHeap) Len() int { return len(h.q) }

func (h *mHeap) Less(i, j int) bool { return h.q[i] > h.q[j] }

func (h *mHeap) Swap(i, j int) {
	h.q[i], h.q[j] = h.q[j], h.q[i]
}

func (h *mHeap) Push(i interface{}) {
	h.q = append(h.q, i.(int))
}

func (h *mHeap) Pop() interface{} {
	x := h.q[len(h.q)-1]
	h.q = h.q[:len(h.q)-1]
	return x
}

// @lc code=start
func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	rs := math.MaxFloat64
	data := make([][]int, len(quality))
	h := &mHeap{}
	// wq := make([]float64, len(quality))

	// 先把工作质量和薪资绑定， 填满data,
	for i := 0; i < len(quality); i++ {
		data[i] = []int{quality[i], wage[i]}
	}
	// 单位工作质量的工资 = wage / quality
	// 按// 单位工作质量的工资 进行排序
	sort.Slice(data, func(i int, j int) bool {
		return float64(data[i][1])/float64(data[i][0]) < float64(data[j][1])/float64(data[j][0])
	})

	sum := 0

	for i := 0; i < len(quality); i++ {
		h.Push(data[i][0])
		sum += data[i][0]

		if h.Len() > k {
			sum -= h.Pop().(int)
		}

		cur := float64(sum) * float64(data[i][0]) / float64(data[i][1])
		if h.Len() == k && cur < rs {
			rs = cur
		}
	}
	// fmt.Println(rs)
	return rs
}

// @lc code=end

