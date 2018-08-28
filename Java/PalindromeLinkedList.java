/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class PalindromeLinkedList {
    public boolean isPalindrome(ListNode head) {
        ListNode slow = head;
        ListNode fast = head;
        ListNode prev = null;
        while(fast != null && fast.next != null) {
            fast = fast.next.next;
            ListNode temp = slow;
            slow = slow.next;
            temp.next = prev;
            prev = temp; 
        }
        if(fast != null) {  
            slow = slow.next;
        } 
        while(prev != null) {
            if (prev.val != slow.val) return false;
            prev = prev.next;
            slow = slow.next;
        }
        return true;
    }
}