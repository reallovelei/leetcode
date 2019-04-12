#include<stdio.h>

define MaxSize 100;
typedef struct BinaryTreeNode
{
    int value;
    BinaryTreeNode* left;
    BinaryTreeNode* right;
}BTree;


void Create(BTree **head, int[] arr);
BTree *Find(BTree *head, int val);

BTree *LchildNode(BTree *head, int val);
BTree *RchildNode(BTree *head, int val);

int Create(BTree **head, int[] arr) {
}
