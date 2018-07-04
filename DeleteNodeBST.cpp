#include <stddef.h>
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

  struct TreeNode {
      int val;
      TreeNode *left;
      TreeNode *right;
      TreeNode(int x) : val(x), left(NULL), right(NULL) {}
  };

class Solution {
public:
    TreeNode* deleteNode(TreeNode*& root, int key) {
            if(root == NULL) {
                return NULL; 
            } else if(key > root->val) {
                root->right = deleteNode(root->right, key);
            } else if(key < root->val) {
                root->left = deleteNode(root->left, key);
            } else {
                if(!root->left) {
                    TreeNode* right = root->right;
                    delete root;
                    root = right;
                } else if(!root->right) {
                    TreeNode* left = root->left;
                    delete root;
                    root = left;
                } else {
                    TreeNode* predecessor = root->left;
                    while(predecessor->right != NULL) {        //find predecessor
                        predecessor = predecessor->right;
                    }
                    root->val = predecessor->val;
                    root->left = deleteNode(root->left, predecessor->val);
                    
                }
            }
            return root;
        }

};