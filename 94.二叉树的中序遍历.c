/*
 * @lc app=leetcode.cn id=94 lang=c
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
 * [94] 二叉树的中序遍历
 * 分析
 * 二叉树中序遍历的特点
 * 对于任意一个节点 先取它的左子节点 再取它本身 最后取它的右节点 如解法一递推公式
 * 【左子节点+根节点+右子节点】
 * 结合栈的特点：先进后出
 * 
 * 题解思路：
 * 1.从根节点 一路向左，遍历左子树，放到栈里,直到最左下角的元素。
 * 2.然后pop 栈顶元素，(如果当前元素有左子树 重复步骤1，将当前节点的左子树顺序push到栈里)
 * 3.如果栈里有元素，就pop 出来。
 */
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
    struct TreeNode data[MAXSIZE];
    int top;
}Stack;

// 初始化栈  申请内存空间  top 赋值
void initStack(Stack *S) {
    S = (Stack*)malloc(sizeof(Stack));
    if (!S) {
        exit(OVERFLOW);
    }
    S->top = 0;
}
bool Push(Stack *S, struct TreeNode* e)
{
    // 栈已满
    if (S->top == MAXSIZE-1) {
        return false;
    }
    S->top++;
    S->data[S->top] = *e;
    return true;
}

struct TreeNode* Pop(Stack *S)
{
    // 栈已空
    if (S->top == 0) {
        return false;
    }
    struct TreeNode* e = (struct TreeNode*)malloc(sizeof(struct TreeNode));
    
    *e = S->data[S->top];
    S->top--;
    
    return e;
}

bool isEmpty(Stack *S) {
    printf("stack top is:%d\n", S->top);
    return S->top == 0 ? true : false;
}

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* inorderTraversal(struct TreeNode* root, int* returnSize){
    //int result[MAXSIZE];
    
    int len = 0;
    int* result = (int*)malloc(MAXSIZE*sizeof(int));   // 存放结果的数组 务必分配内存

    if (root == NULL) {
        *returnSize = len;
        return result;
    }

    struct TreeNode* cur;
    //cur = (struct TreeNode*)malloc(sizeof(struct TreeNode));
    cur = root;

    Stack* s;
    s = (Stack*)malloc(sizeof(Stack));
    initStack(s);
    //int cnt = 0;
    while (cur || !isEmpty(s)) {    // 注意 这里 !isEmpty(s) 这个条件不能放在里面的while,是为了防止 某个子树只有一边有子树的情况。
        //printf("cur val:%d", cur->val);

        while(cur) {    //当前节点不为空  或者栈不为空 // 这里不能加
            if (cur != NULL) {  // 先顺着左子树 一路查找。
                Push(s, cur);
                printf("cur val:%d after push s->top:%d\n", cur->val, s->top);
                cur = cur->left;    // 目的就是要把cur 遍历到null 为止。
            }
        }

        // 如果不为空
        if (!isEmpty(s)) {
            cur = Pop(s);
            printf("after pop cur:%d s->top:%d\n", cur->val, s->top);
            result[len++] = cur->val;

            printf("result:%d \n", cur->val);
            //if (cur->right != NULL) {
                cur = cur->right;
            //}
        }
    }

    *returnSize = len;
    return result;
}



