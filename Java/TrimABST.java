/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class TrimABST {
    public TreeNode trimBST(TreeNode root, int L, int R) {
        if(root != null) {
            
            if (root.val < L) {
                root.left = null;
                return trimBST(root.right, L, R);
            } else if (root.val > R) {
                root.right = null;
                return trimBST(root.left, L, R);
            }
            
            root.left = trimBST(root.left, L, R);
            root.right = trimBST(root.right, L, R); 
        }
        
        return root;
    }
    
}