/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class AddTwoNumbers {
    
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        
        int overflow = 0;
        ListNode head = null;
        ListNode cur = null;
        
        while(l1 != null || l2 != null || overflow > 0) {
            int sum = 0;
            
            if(l1 != null && l2 != null) {
                
                sum = l1.val + l2.val + overflow;
                l1 = l1.next;
                l2 = l2.next;
                
            } else if (l1 != null || l2 != null) {
                
                if(l1 != null) {
                    sum = l1.val + overflow;
                    l1 = l1.next;
                } else {
                    sum = l2.val + overflow;
                    l2 = l2.next;
                }
                
            } else {
                
                sum = overflow;
                
            }
            
            
            int dRsf = sum % 10;
            overflow = sum / 10;
            
            ListNode next = new ListNode(dRsf);
            
            if(cur != null) {
                cur.next = next;
                cur = cur.next;
            } else {
                cur = next;
                head = next;
            }
            
            // if (head == null) {
            //     head = next;
            // }  
        }
        
        return head;
    }
}