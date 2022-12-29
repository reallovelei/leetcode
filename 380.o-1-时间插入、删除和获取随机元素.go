/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] O(1) 时间插入、删除和获取随机元素
 */

// @lc code=start
type RandomizedSet struct {
	nums []int
	hash map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, make(map[int]int, 0)}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.hash[val]
	// 如果有则 直接返回false，不让插入
	if ok {
		return false
	}
	// 注意：下面这2步不能弄反了。
	// 先塞hash,后追加数组
	this.hash[val] = len(this.nums)
	this.nums = append(this.nums, val)

	fmt.Println("insert", val, this.nums, this.hash)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.hash[val]
	// 如果不存在则 直接返回false，不用删除了
	if !ok {
		return false
	}

	tailIndex := len(this.nums) - 1
	// 把要删除的值 对应的下标 ，换成最后一个元素
	this.nums[idx] = this.nums[tailIndex]
	// 把map里的 val => idx 也替换掉， 新的下标

	// insert 1 [0 1] map[0:0 1:1]
	// remove 0 [1] map[1:0] 原来的1 的下标为0了

	this.hash[this.nums[idx]] = idx
	// 把最后一个元素去掉
	this.nums = this.nums[:tailIndex]
	// 删除map
	delete(this.hash, val)

	// fmt.Println("remove", val, this.nums, this.hash)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	// fmt.Println(rand.Intn(len(this.nums)))
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

