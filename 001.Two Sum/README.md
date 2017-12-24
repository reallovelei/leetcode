## 题意
给定一个数组，再给定一个目标值:target。然后在这个数组里找到2个数相加结果是target。然后返回这2个数的下标。
------------------
## 思路
一开始我把这题问我媳妇的时候，我媳妇理解题意后，就说出了2层循环的那个方案。(当然这是我用代码翻译出来的她的方案)
```
func twoSum(nums []int, target int) []int {
	lenNums := len(nums)
	for i := 0; i < lenNums; i++ {
		for j := i + 1; j < lenNums; j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
```
这也是我们第一时间能想到的方案，显然这个方案的时间复杂度是O(n平方) 不是很好。  

经过整理发现可以借用空间一个hash结构。存储map[target - 当前index的value] = 当前index  
一开始我判断是否在map中的时候 是用m[j] 去判断的，提交的时候总是报错。我就看了下才发现 到后面取的时候 当前j 没存进去。
整理后代码如下:
```
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i:=0; i < len(nums); i++ {
        j := target - nums[i]

        // 这里如果是 m[j]  会有问题 因为取的时候 当前j 没存进去
        if _, ok := m[nums[i]]; ok {
            //存在
            return []int {m[nums[i]], i}
            break
        } else {
            fmt.Println(i, nums[i], j)
            m[j] = i
        }
    }
    fmt.Println(m)
    return []int{}
}
```

