/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start
type LRUCache struct {
	size       int           // 当前元素数量
	capacity   int           // 容量
	hash       map[int]*Node // key 对应的node地址
	head, tail *Node         // 双向链表
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

func MakeNode(key int, val int) *Node {
	return &Node{Key: key, Val: val, Prev: nil, Next: nil}
}
func Constructor(capacity int) LRUCache {
	head := MakeNode(-1, -1)
	tail := MakeNode(-1, -1)

	head.Next = tail
	tail.Prev = head
	return LRUCache{
		capacity: capacity,
		hash:     make(map[int]*Node),
		head:     head, tail: tail,
	}
}

func (this *LRUCache) Get(key int) int {

	node, ok := this.hash[key]
	if ok { // 如果是已经存在的，先在链表中删除原有节点, 放到表头。
		this.Remove(node)
		this.AddToHead(node)
		return node.Val
	} else {
		return -1
	}
	// fmt.Println("get", key, this.hash)
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.hash[key]
	if ok { // 如果是已经存在的，先在链表中删除原有节点, 再修改val 放到表头。
		this.Remove(node)
		node.Val = value
		this.AddToHead(node)
	} else {
		node = MakeNode(key, value)
		this.AddToHead(node)
		this.hash[key] = node
		this.size++

		if this.size > this.capacity {
			tailNode := this.tail.Prev
			// fmt.Println("del", this.size, tailNode.Key, this.hash)
			this.Remove(tailNode)
			delete(this.hash, tailNode.Key)
			this.size--
		}
	}
	// fmt.Println("put", key, this.hash)
}

func (this *LRUCache) AddToHead(node *Node) {
	this.head.Next.Prev = node
	node.Next = this.head.Next
	node.Prev = this.head
	this.head.Next = node
}

func (this *LRUCache) Remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

