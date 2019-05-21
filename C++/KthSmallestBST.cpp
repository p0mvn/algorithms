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
    int kthSmallest(TreeNode* root, int k) 
    {
        int res = INT_MIN;
        helper(root, k, res);
        
        return res;
    }
    
    void helper(TreeNode* cur, int& k, int& res)
    {
        if(cur == NULL) return;
        if(res != INT_MIN) return;
        
        helper(cur->left, k, res);
        if(res != INT_MIN) return;
        --k;
        if(k == 0)
        {
            res = cur->val;
            return;
        }
        helper(cur->right, k, res);
    }
    
};