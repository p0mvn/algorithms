/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    int maxDepth(TreeNode* root) 
    {
        return helper(root, 0);
    }
    
    int helper(TreeNode* cur, int depth)
    {
        if(NULL == cur)
        {
            return depth;
        }
        
        int left = helper(cur->left, depth + 1);
        int right = helper(cur->right, depth + 1);
        
        return left > right ? left : right;
    }
};