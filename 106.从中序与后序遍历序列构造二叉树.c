/*
 * @lc app=leetcode.cn id=106 lang=c
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (64.80%)
 * Likes:    125
 * Dislikes: 0
 * Total Accepted:    17.3K
 * Total Submissions: 26.5K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 根据一棵树的中序遍历与后序遍历构造二叉树。
 * 
 * 注意:
 * 你可以假设树中没有重复的元素。
 * 
 * 例如，给出
 * 
 * 中序遍历 inorder = [9,3,15,20,7]
 * 后序遍历 postorder = [9,15,7,20,3]
 * 
 * 返回如下的二叉树：
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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

#include <stdbool.h>
#include <stdio.h>
/*
struct TreeNode* buildSubTree(int* inorder, int* postorder, int startkey, int endkey) {
    struct TreeNode *root, *left, *right;
    int rootval = postorder[endkey]; //找到根节点的val
    
    int rootkey = 0;
    for (;rootkey < endkey; rootkey++) {
        if (inorder[rootkey] == rootval) {
            break;
        }
    }

    //int rootkey = findKeyFromInorder(rootval, endkey, inorder);

    endkey = endkey -1;
    right = buildSubTree(inorder, postorder, rootkey, endkey);
    left = buildSubTree(inorder, postorder, startkey, rootkey);

    root = (struct TreeNode* )malloc(1 * sizeof(struct TreeNode));
    //root = (struct TreeNode *)calloc(1, sizeof(struct TreeNode));
    root->val = rootval;
    root->left = left;
    root->right = right;

    return root;
}
*/

struct TreeNode* buildTree(int* inorder, int inorderSize, int* postorder, int postorderSize) {
    if (inorderSize != postorderSize || inorderSize == 0) {
        return NULL;
    }

    struct TreeNode *root, *left, *right;
    int endkey = inorderSize - 1;
    int rootval = postorder[endkey]; //1.先从后序中 找到根节点的val
    
    
    printf("rootval:%d \n", rootval);
    
    int rootkey = 0;        // 2.再从中序中 查找根节点所在的位置key.
    for (;rootkey < endkey; rootkey++) {
        if (inorder[rootkey] == rootval) {
            break;
        }
    }

    root = (struct TreeNode*)malloc(sizeof(struct TreeNode));
    root->val = rootval;
    root->left = buildTree(inorder, rootkey , postorder, rootkey);
    root->right = buildTree(inorder + rootkey + 1, inorderSize - rootkey - 1, postorder + rootkey, inorderSize - rootkey - 1);

    return root;
}


// @lc code=end

