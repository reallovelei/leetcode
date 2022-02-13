/*
 * @lc app=leetcode.cn id=145 lang=c
 *
 * [145] 二叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Hard (68.19%)
 * Likes:    184
 * Dislikes: 0
 * Total Accepted:    36.7K
 * Total Submissions: 53.5K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 后序 遍历。
 * 
 * 示例:
 * 
 * 输入: [1,null,2,3]  
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3 
 * 
 * 输出: [3,2,1]
 * 
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 * 有个讨巧的做法是按照先序遍历，得到结果再revers一下。
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
    #define MAXSIZE 1000

    // 数组模拟栈结构
    typedef struct Stack
    {
        struct TreeNode* data[MAXSIZE];
        int size;
    }Stack;

    // 初始化栈  申请内存空间  size 赋值
    void initStack(Stack *S) {
        S = (Stack*)malloc(sizeof(Stack));
        if (!S) {
            exit(OVERFLOW);
        }
        S->size = 0;
    }
    bool Push(Stack *S, struct TreeNode* e)
    {
        // 栈已满
        if (S->size == MAXSIZE-1) {
            return false;
        }
        S->size++;
        S->data[S->size] = e;
        return true;
    }

    struct TreeNode* Pop(Stack *S)
    {
        // 栈已空
        if (S->size == 0) {
            return false;
        }
        /*
        struct TreeNode* e = (struct TreeNode*)malloc(sizeof(struct TreeNode));
        e = S->data[S->size];
        S->size--;
        */
        struct TreeNode* e = S->data[S->size--];
        return e;
    }
    /*
    bool isEmpty(Stack *S) {
        printf("stack size is:%d\n", S->size);
        return S->size == 0 ? true : false;
    }
    */
    /**
     * Note: The returned array must be malloced, assume caller calls free().
     */
    int* postorderTraversal(struct TreeNode* root, int* returnSize){

        if (root == NULL) {
            *returnSize = 0;
            return NULL;
        }

        int* result = (int*)malloc(sizeof(int)* MAXSIZE);
        Stack* s = (Stack*)malloc(sizeof(Stack));
        initStack(s);

        struct TreeNode* pcur;  // 当前指针
        struct TreeNode* plast = NULL; // 
        pcur = root;
        int len = 0;

        // 到这里root 就是非空的
        //Push(s, root);
        while (pcur) {  //照样先一路从左子树 走到底。
            Push(s, pcur);
            pcur = pcur->left;
        }

        while(s->size > 0) {    // 这里不能为 >=
            pcur = Pop(s);
            // 判断当前节点的右节点是否可以访问，或者已经访问过了。
            // 如果 不可访问/访问过了  此时可以访问当前节点了。
            if (pcur->right == NULL || pcur->right == plast) {
                result[len++] = pcur->val;
                plast = pcur;   // warning:标记已经到达哪个节点了。
            } else {
                // 如果右节点  没有访问，那么就按照前面的套路将右子树的节点
                //因为当前节点 不能输出2次 所以需要先入栈。
                Push(s, pcur);
                pcur = pcur->right;

                while (pcur) {  //照样先一路从左子树 走到底。
                    Push(s, pcur);
                    pcur = pcur->left;
                }
            }
        }


        *returnSize = len;
        return result;
    }


// @lc code=end

