题目:求众数
给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在众数。
尽量不要使用空间复杂度。

### 示例 1:
输入: [3,2,3]
输出: 3

### 示例 2:
输入: [2,2,1,1,1,2,2]
输出: 2

### 解题思路
一开始我是想用一个map结构存储每个数出现的次数。最后取一个最大的,当然也可以每次set的时候保留最大次数的那个。
后来发现不让占用空间复杂度。那就只能记一个数了。
于是先认为第一个数就是众数，如果下一个和前一个数一样，那么count +1,如果不一样 count 如果>0,  count -1. 如果不大于0 代表这个数肯定不是众数 则重新计数。
