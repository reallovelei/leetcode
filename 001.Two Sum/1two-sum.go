package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i:=0; i < len(nums); i++ {
		/*
		if (nums[j] > target) {
            j--
            continue
        }
		*/

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

func main() {
	// bad case
	// sorted 后索引值 会改变。所以不太可能
	// {3,2,4}    6

	nums := []int {3,2, 4}
	twoSum(nums, 6)
}
