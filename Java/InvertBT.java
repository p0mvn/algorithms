/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class InvertBT {
    public TreeNode invertTree(TreeNode root) {
        
        if (root == null) {
            return root;
        } 
        
        TreeNode temp = root.left;
        root.left = root.right;
        root.right = temp;
        root.right = invertTree(root.right);
        root.left = invertTree(root.left);
        
        return root;
        
    }
}