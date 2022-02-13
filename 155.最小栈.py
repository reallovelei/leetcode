#
# @lc app=leetcode.cn id=155 lang=python
#
# [155] 最小栈
#
# https://leetcode-cn.com/problems/min-stack/description/
#
# algorithms
# Easy (50.33%)
# Likes:    308
# Dislikes: 0
# Total Accepted:    52.4K
# Total Submissions: 104K
# Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]'
#
# 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
# 
# 
# push(x) -- 将元素 x 推入栈中。
# pop() -- 删除栈顶的元素。
# top() -- 获取栈顶元素。
# getMin() -- 检索栈中的最小元素。
# 
# 
# 示例:
# 
# MinStack minStack = new MinStack();
# minStack.push(-2);
# minStack.push(0);
# minStack.push(-3);
# minStack.getMin();   --> 返回 -3.
# minStack.pop();
# minStack.top();      --> 返回 0.
# minStack.getMin();   --> 返回 -2.
# 
# 
#

# @lc code=start
class MinStack(object):

    def __init__(self):
        """
        initialize your data structure here.
        """
        # 数据栈
        self.data = []
        # 辅助栈
        self.helper = []


    def push(self, x):
        """
        :type x: int
        :rtype: None
        """
        #数据栈 先插入
        self.data.append(x)
        #辅助栈 如果没有数据，或者 x < 辅助栈 栈顶元素(当前最小元素)
        if len(self.helper) == 0 or x <= self.helper[-1]:
            self.helper.append(x)
        else:
            self.helper.append(self.helper[-1])
        

    def pop(self):
        """
        :rtype: None
        """
        # 如果 data还有数据，先pop 辅助栈，在pop数据栈。
        # 如果数据栈没有数据 就不返回了？
        if self.data:
            self.helper.pop()
            return self.data.pop()
        

    def top(self):
        """
        :rtype: int
        """
        if self.data:
            return self.data[-1]
        

    def getMin(self):
        """
        :rtype: int
        """
        if self.helper:
            return self.helper[-1]
        


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(x)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()
# @lc code=end

