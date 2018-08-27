/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class OddEvenLinkedList {
    public ListNode oddEvenList(ListNode head) {
        if(head == null || head.next == null) return head;
        
        ListNode oddHead = new ListNode(-1);
        
        oddHead.next = head;
        ListNode odd = oddHead.next;
        ListNode even  = odd.next;
        ListNode evenHead = new ListNode(-1);
        evenHead.next = even;
        
        
        while(even != null && even.next != null) {
            
            odd.next = even.next;
            if(odd.next != null) {
                even.next = odd.next.next;
            }
            
            
            odd = odd.next;
            even = even.next;
        }
        odd.next = evenHead.next;
        
        return oddHead.next;
        
        
        
    }
}