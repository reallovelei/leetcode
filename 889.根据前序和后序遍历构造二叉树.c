/*
 * @lc app=leetcode.cn id=889 lang=c
 *
 * [889] 根据前序和后序遍历构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (58.93%)
 * Likes:    43
 * Dislikes: 0
 * Total Accepted:    2.3K
 * Total Submissions: 3.9K
 * Testcase Example:  '[1,2,4,5,3,6,7]\n[4,5,2,6,7,3,1]'
 *
 * 返回与给定的前序和后序遍历匹配的任何二叉树。
 * 
 * pre 和 post 遍历中的值是不同的正整数。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
 * 输出：[1,2,3,4,5,6,7]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= pre.length == post.length <= 30
 * pre[] 和 post[] 都是 1, 2, ..., pre.length 的排列
 * 每个输入保证至少有一个答案。如果有多个答案，可以返回其中一个。
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */

#include <stdio.h>
#include <stdbool.h>

struct TreeNode* constructFromPrePost(int* pre, int preSize, int* post, int postSize){
    if (preorderSize != inorderSize || preorderSize <= 0) { //注意: 这里要加 = 不然为0的时候还继续往下执行，会出现数组下标 -1的情况
        return NULL;
    }

    int root = pre[0];
    int leftRootKey, rightRootKey, leftCnt, rightCnt = 0; // 左子树根节点/ 右子树根节点/ 左子树个数 / 右子树个数
    



}


// @lc code=end

