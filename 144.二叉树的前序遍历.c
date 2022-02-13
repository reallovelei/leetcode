/*
 * @lc app=leetcode.cn id=144 lang=c
 *
 * [144] 二叉树的前序遍历
 */
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAXSIZE 10000
#define OVERFLOW 0
#define error -65530

/**栈的数据结构定义 用数组模拟的栈 **/
typedef struct stack
{
    struct TreeNode data[MAXSIZE];  // 栈里的元素数组
    int top;        //  栈顶指针
}sqstack;

// 初始化栈  申请内存空间  top 赋值
void initStack(sqstack *S) {
    S = (sqstack*)malloc(sizeof(sqstack));
    if (!S) {
        exit(OVERFLOW);
    }
    S->top = 0;
}
bool Push(sqstack *S, struct TreeNode* e)
{
    // 栈已满
    if (S->top == MAXSIZE-1) {
        return false;
    }
    S->top++;
    S->data[S->top] = *e;    // 顺序??
    
    return true;
}

struct TreeNode* Pop(sqstack *S)
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

bool isEmpty(sqstack *S) {
    return S->top == 0 ? true : false;
}


int* preorderTraversal(struct TreeNode* root, int* returnSize){
    int len = 0;    // 元素个数
    if (root == NULL) {
        *returnSize = len;    // @todo 注意:这个size 一定要赋值,不然 [] 这个case会报错:load of null pointer of type 'int'
        return NULL;
    }
    
    sqstack* s = (sqstack*)malloc(sizeof(sqstack));
    initStack(s);
    
    int* result = (int*)malloc(1000*sizeof(int));   // 存放结果的数组
    
    while (root || !isEmpty(s)) {
        while(root) {   // 将树的左子树全部入栈。直到最左下的子树为空。
        
            result[len++] = root->val;
            Push(s, root);  // 将当前节点(左节点)入栈，
            root = root->left;
        }
        // 当前节点的 左子树 都遍历完了。

        if (!isEmpty(s)) {
            root = Pop(s);  // 弹出栈  找右孩子节点。  
            if (root != NULL) {
                root = root->right;
            }
        }
    }
    *returnSize = len;
    return result;

}

/**
 * Note: The returned array must be malloced, assume caller calls free().
 * 先序遍历是先输出根节点，再输出左子树，最后输出右子树
 
int* preorderTraversal(struct TreeNode* root, int* returnSize){
    if (root == NULL) return NULL;

    struct TreeNode* stack = (struct TreeNode*)malloc(sizeof(struct TreeNode*)* 1000);
    struct TreeNode* tmp = root;         // 根指针
    int top = 0; // 栈顶指针
    int size = 0;
    int* res = (int*)malloc(sizeof(int) * 1000);
    //int *res = (int*)malloc(sizeof(int) * *returnSize);
    // 左边不断进栈，
    while (tmp || top != 0) {
        while (tmp != NULL) {
            stack[top++] = tmp; // 根节点先进
            tmp = tmp->left;    // 访问左子树
        }

        if (top != 0) {
            tmp = stack[top]; //开始出栈
            if (tmp != NULL) {
                res[size++] = tmp->val;  // 扔到数组里
                tmp = tmp->right;
            }
            top--;
        }
    }
    return res;
}
*/


