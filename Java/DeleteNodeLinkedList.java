/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class DeleteNodeLinkedList {
    public void deleteNode(ListNode node) {
        if(node == null) return;
        if(node.next == null) return;
        node.val = node.next.val;
        node.next = node.next.next;
    }
}