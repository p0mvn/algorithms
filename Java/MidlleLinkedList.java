/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class MiddleLinkedList {

    public ListNode middleNodeVersionTwo(ListNode head) {
        
        ListNode slow = head;
        ListNode fast = head;
        while(fast != null && fast.next != null) {
            fast = fast.next.next;
            slow = slow.next;
        }
        
        return slow;
        
    }


    public ListNode middleNode(ListNode head) {
        
        return helper(head, 1, head);
        
    }
    
    private ListNode helper(ListNode cur, int count, ListNode mid) {
        
        
        
        if (cur == null) {
            return mid;
        }
        
        if(count % 2 == 0) {
            mid = mid.next;
        }
        
        return helper(cur.next, ++count, mid);
        
        
        
    }
}