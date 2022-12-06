/*
 * @lc app=leetcode.cn id=895 lang=golang
 *
 * [895] 最大频率栈
 */

// @lc code=start
type FreqStack struct {
	cnt    map[int]int // 记录每个数有几次
	stacks [][]int     // 多个栈   按次数算的。
}

func Constructor() FreqStack {
	// 初始化map
	return FreqStack{cnt: make(map[int]int)}
}

func (this *FreqStack) Push(val int) {
	c := this.cnt[val] // 先拿到入栈元素 的数量
	if c == len(this.stacks) {
		this.stacks = append(this.stacks, []int{val})
	} else {
		this.stacks[c] = append(this.stacks[c], val)
	}

	// 更新计数
	this.cnt[val]++

}

func (this *FreqStack) Pop() int {
	lastStackIdx := len(this.stacks) - 1
	lastStack := this.stacks[lastStackIdx]
	lastKey := len(lastStack) - 1

	val := lastStack[lastKey]
	lastStack = lastStack[:lastKey]

	if len(lastStack) == 0 {
		this.stacks = this.stacks[:lastStackIdx]
	} else {
		this.stacks[lastStackIdx] = lastStack
	}
	this.cnt[val]--

	return val

}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
// @lc code=end

