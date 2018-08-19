/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class ReverseLinkedListII {
    public ListNode reverseBetween(ListNode head, int m, int n) {
        
        int count = 1;
        ListNode prev = null;
        ListNode cur = head;
        ListNode post;
        
        while(count < m) {
            prev = cur;
            cur = cur.next;
            count++;
        }
        post = cur.next;
        
        
        for(int i = count; i <= n; i++) {
            cur.next = prev;
            prev = cur;
            cur = post;
            post = post.next;
        }
        
        return head;
    }
    
}