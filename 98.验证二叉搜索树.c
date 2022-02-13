/*
 * @lc app=leetcode.cn id=98 lang=c
 *
 * [98] 验证二叉搜索树
 */
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */

// 找到左子树最大的值 感觉每次都找一遍，有点浪费 啊
void findLeftMax(struct TreeNode* left, int *val) {
    if (!left) {
        return ;
    }

    if (left->val > *val) {
        *val = left->val;
    }

    findLeftMax(left->left, val);
    findLeftMax(left->right, val);
}

// 遍历右边找出最小的值。
void findRightMin(struct TreeNode* right, int *val) {
    if (!right) {
        return;
    }

    if (right->val < *val) {
        *val = right->val;
    }

    findRightMin(right->left, val);
    findRightMin(right->right, val);
}

bool isValidBST(struct TreeNode* root){
    if (!root) return true;
    int max = INT_MIN;
    int min = INT_MAX;
    
    findLeftMax(root->left, &max);
    findRightMin(root->right, &min);
    //printf("root = %d, leftmax = %d, rightmin = %d \n",root->val, max, min);
    bool ret = true;
    
    if (root->left != NULL ) {
        ret = ret && isValidBST(root->left)  && root->val > max;
        if (!root->right)
        return ret;
    }
    

    if (root->right != NULL)
    {
        ret = ret && isValidBST(root->right) && root->val < min;
        if (!root->left)
        return ret;
    }
    return ret;
}



