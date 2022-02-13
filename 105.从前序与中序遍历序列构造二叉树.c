/*
 * @lc app=leetcode.cn id=105 lang=c
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (61.98%)
 * Likes:    274
 * Dislikes: 0
 * Total Accepted:    31K
 * Total Submissions: 49.5K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
 * 
 * 注意:
 * 你可以假设树中没有重复的元素。
 * 
 * 例如，给出
 * 
 * 前序遍历 preorder = [3,9,20,15,7]
 * 中序遍历 inorder = [9,3,15,20,7]
 * 
 * 返回如下的二叉树：
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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

struct TreeNode* buildTree(int* preorder, int preorderSize, int* inorder, int inorderSize){
    if (preorderSize != inorderSize || preorderSize <= 0) { //注意: 这里要加 = 不然为0的时候还继续往下执行，会出现数组下标 -1的情况
        return NULL;
    }
    struct TreeNode* root, left, right;
    root = (struct TreeNode*)malloc(sizeof(struct TreeNode));

    int rootval = preorder[0];  // 从前序序列中找到 当前的根节点。
    //printf("\nrootval:%d\n", rootval);

    int rootkey = 0;
    // start 找到根节点 在中序 序列中的位置。 注意:这里size一定要-1 不然数组越界
    for (;rootkey < preorderSize-1; rootkey++) {
        //printf("curindex:%d curval:%d", rootkey, inorder[rootkey]);
        if (inorder[rootkey] == rootval) {
            break;
        }
    }
    // end

    int leftCnt = rootkey;
    int rightCnt = preorderSize-rootkey-1;
    //printf("\n rootkey:%d preorderSize:%d inorderSize:%d leftcnt:%d rightcnt:%d\n", rootkey,preorderSize, inorderSize, leftCnt, rightCnt);
    

    root->val = rootval;
    root->left = buildTree(preorder+1, leftCnt, inorder, leftCnt);
    root->right = buildTree(preorder+rootkey+1, rightCnt, inorder+rootkey+1, rightCnt);

    return root;
}


// @lc code=end

