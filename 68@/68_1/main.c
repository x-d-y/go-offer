#include "stdio.h"

int main()
{
}

struct TreeNode* lowestCommonAncestor(struct TreeNode* root, struct TreeNode* p, struct TreeNode* q) {
    if (p.val < root.val && q.val < root.val ){
        return lowestCommonAncestor(root->left, p, q)
    }else if (p.val > root.val && q.val > root.val){
        return lowestCommonAncestor(root->right, p, q)
    }else {
        return root
    }
}

struct TreeNode {
      int val;
      struct TreeNode *left;
      struct TreeNode *right;
};