/*
 * @lc app=leetcode.cn id=1801 lang=golang
 *
 * [1801] 积压订单中的订单总数
 */

// @lc code=start
func getNumberOfBacklogOrders(orders [][]int) int {
	buy, sell := buyOrders{}, sellOrders{}
	for _, o := range orders {
		price, amount, oType := o[0], o[1], o[2]

		if oType == 0 { // buy order
			// 当前订单数量 > 0, 有待售卖订单，且最低价格比当前购买订单的价格低 或者等于
			for amount > 0 && len(sell) > 0 && sell[0].price <= price {
				// 第一个 待售卖订单的数量 > 当前订单数量。
				if sell[0].cnt > amount {
					// 撮合成交，把售卖订单的数量 减掉
					sell[0].cnt -= amount
					amount = 0
					break
				}
				// 第一个售卖订单 吃不下我这个大单。
				// 相当于把第一个 sell订单 吃掉了。
				amount -= heap.Pop(&sell).(order).cnt // heap操作的是指针
			}
			// 如果把sell订单都找完了 都没有吃下我这个大单. 就要往buy里放了。
			if amount > 0 {
				heap.Push(&buy, order{price, amount}) // 这里的amount 可能已经成交了一部分了，所以可能小于 原来的amount
			}
		} else { // sell
			for amount > 0 && len(buy) > 0 && buy[0].price >= price {
				if buy[0].cnt > amount {
					buy[0].cnt -= amount
					amount = 0
					break
				}
				amount -= heap.Pop(&buy).(order).cnt
			}

			if amount > 0 {
				heap.Push(&sell, order{price, amount})
			}
		}
	}

	ans := 0
	for _, o := range buy {
		ans += o.cnt
	}

	for _, o := range sell {
		ans += o.cnt
	}

	return ans % (1e9 + 7)
}

type order struct {
	price int
	cnt   int
}

type buyOrders []order // 定义积压订单
// 堆操作
func (buy buyOrders) Len() int { return len(buy) }

// 购买积压订单 要从大到小排序 大顶堆 (价高者得，找冤大头)
func (buy buyOrders) Less(i, j int) bool  { return buy[i].price > buy[j].price }
func (buy buyOrders) Swap(i, j int)       { buy[i], buy[j] = buy[j], buy[i] }
func (buy *buyOrders) Push(v interface{}) { *buy = append(*buy, v.(order)) }
func (buy *buyOrders) Pop() interface{} {
	tmp := *buy
	node := tmp[len(tmp)-1] // 拿到最后一个节点
	*buy = tmp[:len(tmp)-1]
	return node
}

//  售卖订单
type sellOrders []order          // 定义积压订单
func (sell sellOrders) Len() int { return len(sell) }

// 售卖积压订单 要从小到大排序 小顶堆(尽量买便宜的)
func (sell sellOrders) Less(i, j int) bool  { return sell[i].price < sell[j].price }
func (sell sellOrders) Swap(i, j int)       { sell[i], sell[j] = sell[j], sell[i] }
func (sell *sellOrders) Push(v interface{}) { *sell = append(*sell, v.(order)) }
func (sell *sellOrders) Pop() interface{} {
	tmp := *sell
	node := tmp[len(tmp)-1] // 拿到最后一个节点
	*sell = tmp[:len(tmp)-1]
	return node
}

// @lc code=end
