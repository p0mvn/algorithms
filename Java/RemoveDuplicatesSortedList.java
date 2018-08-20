/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class RemoveDuplicatesSortedList {
    
    public ListNode deleteDuplicates(ListNode head) {
        if(head == null) return head;
        
        ListNode cur = head;
        
        while(cur != null && cur.next != null) {
            if(cur.next.val == cur.val) {
                cur.next = cur.next.next;
            } else {
                cur = cur.next;

            }
        }

        return head;
        
    }
    
    
    
    public ListNode deleteDuplicatesOne(ListNode head) {
        if(head == null) return head;
        
        Set seen = new HashSet();
        ListNode prev = head;
        seen.add(prev.val);
        ListNode cur = prev.next;
        while(cur != null) {
            if(seen.contains(cur.val)) {
                prev.next = cur.next;
            } else {
                seen.add(cur.val);
                prev = prev.next;
            }
            cur = cur.next;
        }
        
        return head;
        
    }
}